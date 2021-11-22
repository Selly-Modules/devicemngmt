package devicemngmt

import (
	"context"

	"github.com/Selly-Modules/logger"
	"github.com/Selly-Modules/mongodb"
	"go.mongodb.org/mongo-driver/bson"
)

// FindAllDevicesByUserID ...
func (s Service) FindAllDevicesByUserID(userID string) []Device {
	var (
		ctx    = context.Background()
		col    = s.getDeviceCollection()
		result = make([]Device, 0)
		id, _  = mongodb.NewIDFromString(userID)
		cond   = bson.M{
			"userId": id,
		}
	)

	// Find
	cursor, err := col.Find(ctx, cond)
	if err != nil {
		logger.Error("devicemngmt - findAllDevicesByUserID ", logger.LogData{
			"err": err.Error(),
		})
		return result
	}
	defer cursor.Close(ctx)
	cursor.All(ctx, &result)

	return result
}
