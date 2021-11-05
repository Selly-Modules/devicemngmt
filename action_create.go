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
	deviceData, err := payload.newDevice()
	if err != nil {
		return err
	}

	// Find deviceID existed or not
	if s.isDeviceIDExisted(ctx, deviceData.DeviceID) {
		return errors.New("this device is already existed")
	}

	// Create device
	_, err = col.InsertOne(ctx, deviceData)
	if err != nil {
		logger.Error("devicemngt - Create ", logger.LogData{
			"doc": deviceData,
			"err": err.Error(),
		})
		return fmt.Errorf("error when create device: %s", err.Error())
	}

	return nil
}

func (payload CreateOptions) newDevice() (result Device, err error) {
	timeNow := now()
	device := Device{
		ID:              mongodb.NewObjectID(),
		DeviceID:        payload.DeviceID,
		OSName:          getOSName(payload.UserAgent),
		OSVersion:       getOSVersion(payload.UserAgent),
		IP:              payload.IP,
		Language:        getLanguage(payload.Language),
		AuthToken:       payload.AuthToken,
		LastActivatedAt: timeNow,
		CreatedAt:       timeNow,
		FCMToken:        payload.FCMToken,
		Model:           payload.Model,
		Manufacturer:    payload.Manufacturer,
	}

	// App version
	if payload.AppVersion != "" {
		device.AppVersion = payload.AppVersion
		device.IsMobile = true
	}

	result = device
	return
}
