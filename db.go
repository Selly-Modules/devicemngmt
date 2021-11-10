package devicemngmt

import (
	"context"
	"fmt"

	"github.com/Selly-Modules/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

//  getDeviceCollection ...
func (s Service) getDeviceCollection() *mongo.Collection {
	if s.TablePrefix != "" {
		return s.DB.Collection(fmt.Sprintf("%s-%s", s.TablePrefix, tableDevice))
	}
	return s.DB.Collection(tableDevice)
}

func (s Service) isDeviceIDExisted(ctx context.Context, deviceID string) bool {
	var (
		col    = s.getDeviceCollection()
		device = Device{}
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
