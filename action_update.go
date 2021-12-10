package devicemngmt

import (
	"context"
	"errors"
	"fmt"

	"github.com/Selly-Modules/logger"
	"github.com/Selly-Modules/mongodb"
	"go.mongodb.org/mongo-driver/bson"
)

// UpdateOptions ...
type UpdateOptions struct {
	UserID       string
	UserAgent    string
	AppVersion   string
	IP           string
	FCMToken     string
	AuthToken    string
	Language     string
	Model        string
	Manufacturer string
	OsName       string
	OsVersion    string
	IsMobile     bool
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
	if osName == "" || osVersion == "" {
		osName = payload.OsName
		osVersion = payload.OsVersion
		isMobile = payload.IsMobile
	}

	// Setup update data
	userID, _ := mongodb.NewIDFromString(payload.UserID)
	updateData := bson.M{
		"$set": bson.M{
			"userId":          userID,
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
