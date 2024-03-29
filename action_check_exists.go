package devicemngmt

import (
	"context"

	"github.com/Selly-Modules/logger"
	"go.mongodb.org/mongo-driver/bson"
)

func (s Service) IsDeviceIDExisted(deviceID string) bool {
	var (
		col = s.getDeviceCollection()
		ctx = context.Background()
	)

	total, err := col.CountDocuments(ctx, bson.M{"deviceId": deviceID})
	if err != nil {
		logger.Error("devicemngmt - isDeviceIDExisted", logger.LogData{
			Source:  "devicemngmt.IsDeviceIDExisted",
			Message: err.Error(),
			Data:    deviceID,
		})
		return true
	}
	return total != 0
}
