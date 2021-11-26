package devicemngmt

import (
	"context"

	"github.com/Selly-Modules/logger"
	"go.mongodb.org/mongo-driver/bson"
)

// GetUserIDByAuthToken ...
func (s Service) GetUserIDByAuthToken(authToken string) (userID string) {
	var (
		ctx    = context.Background()
		col    = s.getDeviceCollection()
		device = Device{}
		cond   = bson.M{
			"authToken": authToken,
		}
	)

	if authToken == "" {
		return
	}

	// Find
	if err := col.FindOne(ctx, cond).Decode(&device); err != nil {
		logger.Error("devicemngmt - getUserIDByAuthToken", logger.LogData{
			"authToken": authToken,
			"err":       err.Error(),
		})
		return
	}

	userID = device.UserID.Hex()
	return
}

// GetUserIDByDeviceId ...
func (s Service) GetUserIDByDeviceId(deviceId string) (userID string) {
	var (
		ctx    = context.Background()
		col    = s.getDeviceCollection()
		device = Device{}
		cond   = bson.M{
			"deviceId": deviceId,
		}
	)

	if deviceId == "" {
		return
	}

	// Find
	if err := col.FindOne(ctx, cond).Decode(&device); err != nil {
		logger.Error("devicemngmt - getUserIDByDeviceId", logger.LogData{
			"deviceId": deviceId,
			"err":      err.Error(),
		})
		return
	}

	userID = device.UserID.Hex()
	return
}
