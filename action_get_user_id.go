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
			Source:  "devicemngmt.GetUserIDByAuthToken",
			Message: err.Error(),
			Data:    authToken,
		})
		return
	}

	userID = device.UserID.Hex()
	return
}
