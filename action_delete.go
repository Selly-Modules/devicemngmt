package devicemngmt

import (
	"context"
	"fmt"

	"github.com/Selly-Modules/logger"
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

	// Delete
	if _, err := col.DeleteOne(ctx, cond); err != nil {
		logger.Error("devicemngt - deleteByDeviceID", logger.LogData{
			"deviceId": deviceID,
			"err":      err.Error(),
		})
		return fmt.Errorf("error when delete device: %s", err.Error())
	}

	return nil
}
