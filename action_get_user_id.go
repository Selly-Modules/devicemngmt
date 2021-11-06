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

	// Find
	if err := col.FindOne(ctx, cond).Decode(&device); err != nil {
		logger.Error("devicemngt - getUserIDByAuthToken", logger.LogData{
			"authToken": authToken,
			"err":       err.Error(),
		})
		return
	}

	userID = device.UserID
	return
}