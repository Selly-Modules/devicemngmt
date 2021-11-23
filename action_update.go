package devicemngmt

import (
	"context"
	"errors"
	"fmt"

	"github.com/Selly-Modules/logger"
	"go.mongodb.org/mongo-driver/bson"
)

// UpdateOptions ...
type UpdateOptions struct {
	UserAgent    string
	AppVersion   string
	IP           string
	FCMToken     string
	AuthToken    string
	Language     string
	Model        string
	Manufacturer string
}

// UpdateByDeviceID ...
func (s Service) UpdateByDeviceID(deviceID string, payload UpdateOptions) error {
	var (
		ctx  = context.Background()
		col  = s.getDeviceCollection()
		cond = bson.M{
			"deviceId": deviceID,
		}
	)

	total, _ := col.CountDocuments(ctx, cond)
	if total == 0 {
		return errors.New("deviceId not found")
	}

	// Validate payload
	err := payload.validate()
	if err != nil {
		return err
	}

	// Get userAgent data
	osName, osVersion, isMobile := getUserAgentData(payload.UserAgent)

	// Setup update data
	updateData := bson.M{
		"$set": bson.M{
			"osName":          osName,
			"osVersion":       osVersion,
			"ip":              payload.IP,
			"language":        getLanguage(payload.Language),
			"authToken":       payload.AuthToken,
			"fcmToken":        payload.FCMToken,
			"model":           payload.Model,
			"manufacturer":    payload.Manufacturer,
			"appVersion":      payload.AppVersion,
			"isMobile":        isMobile,
			"lastActivatedAt": now(),
		},
	}

	// Update
	_, err = col.UpdateOne(ctx, cond, updateData)
	if err != nil {
		logger.Error("devicemngmt - updateByDeviceID", logger.LogData{
			"deviceId": deviceID,
			"err":      err.Error(),
		})
		return fmt.Errorf("error when update device: %s", err.Error())
	}

	return nil
}
