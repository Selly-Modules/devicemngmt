package devicemngmt

import (
	"context"

	"github.com/Selly-Modules/logger"
	"go.mongodb.org/mongo-driver/bson"
)

func (s Service) IsDeviceIDExisted(deviceID string) bool {
	var (
		col    = s.getDeviceCollection()
		device = Device{}
		ctx    = context.Background()
	)

	if err := col.FindOne(ctx, bson.M{"deviceId": deviceID}).Decode(&device); err != nil {
		logger.Error("devicemngmt - findByDeviceID", logger.LogData{
			"deviceId": deviceID,
			"err":      err.Error(),
		})
		return true
	}
	return !device.ID.IsZero()
}
