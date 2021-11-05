package devicemngmt

import (
	"context"
	"errors"

	"github.com/Selly-Modules/logger"
	"github.com/Selly-Modules/mongodb"
	ua "github.com/mssola/user_agent"
	"go.mongodb.org/mongo-driver/bson"
)

// DeviceCreate ...
type DeviceCreate struct {
	DeviceID   string
	UserID     string
	UserAgent  string
	AppVersion string
	IP         string
	FCMToken   string
	Language   string
}

// Create ...
func (s Service) Create(payload DeviceCreate) error {
	var (
		col = s.getDeviceCollection()
		ctx = context.Background()
	)

	// New device data from payload
	deviceData, err := payload.newDevice()
	if err != nil {
		return err
	}

	// Find device id existed or not
	device := Device{}
	if err = col.FindOne(ctx, bson.M{"deviceID": deviceData.DeviceID}).Decode(&device); err != nil {
		logger.Error("devicemngt - findByDeviceID", logger.LogData{
			"deviceID": deviceData.DeviceID,
			"err":      err.Error(),
		})
	}
	if !device.ID.IsZero() {
		return errors.New("deviceID already exists")
	}

	// Create device
	_, err = col.InsertOne(ctx, deviceData)
	if err != nil {
		logger.Error("devicemngt - Create ", logger.LogData{
			"doc": deviceData,
			"err": err.Error(),
		})
		return errors.New("create device fail")
	}

	return nil
}

func (payload DeviceCreate) newDevice() (result Device, err error) {
	timeNow := now()
	device := Device{
		ID:             mongodb.NewObjectID(),
		LastActivityAt: timeNow,
		CreatedAt:      timeNow,
		FCMToken:       payload.FCMToken,
	}

	// Set deviceID
	if payload.DeviceID == "" {
		logger.Error("devicemngt - Create: no deviceID data", logger.LogData{
			"payload": payload,
		})
		err = errors.New("no deviceID data")
		return
	}
	device.DeviceID = payload.DeviceID

	// 	OSName, OSVersion
	if payload.UserAgent == "" {
		logger.Error("devicemngt - Create: no userAgent data", logger.LogData{
			"payload": payload,
		})
		err = errors.New("no userAgent data")
		return
	}
	uaData := ua.New(payload.UserAgent)
	device.OSName = uaData.OSInfo().Name
	device.OSVersion = uaData.OSInfo().Version

	// App version
	if payload.AppVersion != "" {
		device.AppVersion = payload.AppVersion
		device.IsMobile = true
	}

	// IP
	if payload.IP == "" {
		logger.Error("devicemngt - Create: no ip data", logger.LogData{
			"payload": payload,
		})
		err = errors.New("no ip data")
		return
	}

	// Language, default is vietnamese(vi)
	if payload.Language == "" {
		device.Language = viLanguage
	} else {
		device.Language = enLanguage
	}

	// userIDe
	userID, _ := mongodb.NewIDFromString(payload.UserID)
	if userID.IsZero() {
		logger.Error("devicemngt - Create: invalid userID data", logger.LogData{
			"payload": payload,
		})
		err = errors.New("invalid userID data")
		return
	}

	// Generate authToken from userID
	device.AuthToken = s.generateAuthToken(userID)

	result = device
	return
}
