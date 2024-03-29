package devicemngmt

import (
	"context"
	"errors"
	"fmt"

	"github.com/Selly-Modules/logger"
	"github.com/Selly-Modules/mongodb"
	"go.mongodb.org/mongo-driver/bson"
)

// DeleteByDeviceID ...
func (s Service) DeleteByDeviceID(deviceID string) error {
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

	// Delete
	if _, err := col.DeleteOne(ctx, cond); err != nil {
		logger.Error("devicemngmt - deleteByDeviceID", logger.LogData{
			Source:  "devicemngmt.DeleteByDeviceID",
			Message: err.Error(),
			Data:    deviceID,
		})
		return fmt.Errorf("error when delete device: %s", err.Error())
	}

	return nil
}

// DeleteByUserID ...
func (s Service) DeleteByUserID(userID string) error {
	var (
		ctx = context.Background()
		col = s.getDeviceCollection()
	)

	id, isValid := mongodb.NewIDFromString(userID)
	if !isValid {
		return errors.New("invalid userID data")
	}

	cond := bson.M{
		"userId": id,
	}

	// Delete
	if _, err := col.DeleteMany(ctx, cond); err != nil {
		logger.Error("devicemngmt - deleteByUserID", logger.LogData{
			Source:  "devicemngmt.DeleteByUserID",
			Message: err.Error(),
			Data:    userID,
		})
		return fmt.Errorf("error when delete device by userId: %s", err.Error())
	}

	return nil
}
