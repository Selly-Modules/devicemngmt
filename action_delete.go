package devicemngmt

import (
	"context"
	"fmt"

	"github.com/Selly-Modules/logger"
	"go.mongodb.org/mongo-driver/bson"
)

// DeleteDeviceByDeviceID ...
func (s Service) DeleteDeviceByDeviceID(deviceID string) error {
	var (
		ctx  = context.Background()
		col  = s.getDeviceCollection()
		cond = bson.M{
			"deviceID": deviceID,
		}
	)

	// Delete
	if _, err := col.DeleteOne(ctx, cond); err != nil {
		logger.Error("devicemngt - deleteDeviceByDeviceID", logger.LogData{
			"deviceID": deviceID,
			"err":      err.Error(),
		})
		return fmt.Errorf("error when delete device: %s", err.Error())
	}

	return nil
}
