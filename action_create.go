package devicemngmt

import (
	"context"
	"errors"
	"fmt"

	"github.com/Selly-Modules/logger"
	"github.com/Selly-Modules/mongodb"
)

// CreateOptions ...
type CreateOptions struct {
	DeviceID     string
	UserID       string
	UserAgent    string
	AppVersion   string
	IP           string
	FCMToken     string
	AuthToken    string
	Language     string
	Model        string
	Manufacturer string
}

// Create ...
func (s Service) Create(payload CreateOptions) error {
	var (
		col = s.getDeviceCollection()
		ctx = context.Background()
	)

	// Validate payload
	err := payload.validate()
	if err != nil {
		return err
	}

	// New device data from payload
	deviceData := payload.newDevice()

	// Find deviceID existed or not
	if s.IsDeviceIDExisted(ctx, deviceData.DeviceID) {
		return errors.New("this device is already existed")
	}

	// Create device
	_, err = col.InsertOne(ctx, deviceData)
	if err != nil {
		logger.Error("devicemngmt - Create ", logger.LogData{
			"doc": deviceData,
			"err": err.Error(),
		})
		return fmt.Errorf("error when create device: %s", err.Error())
	}

	return nil
}

func (payload CreateOptions) newDevice() Device {
	timeNow := now()

	// Get userAgent data
	osName, osVersion, isMobile := getUserAgentData(payload.UserAgent)

	userID, _ := mongodb.NewIDFromString(payload.UserID)
	return Device{
		ID:              mongodb.NewObjectID(),
		DeviceID:        payload.DeviceID,
		OSName:          osName,
		OSVersion:       osVersion,
		IP:              payload.IP,
		Language:        getLanguage(payload.Language),
		AuthToken:       payload.AuthToken,
		LastActivatedAt: timeNow,
		CreatedAt:       timeNow,
		FCMToken:        payload.FCMToken,
		Model:           payload.Model,
		Manufacturer:    payload.Manufacturer,
		UserID:          userID,
		IsMobile:        isMobile,
		AppVersion:      payload.AppVersion,
	}
}
