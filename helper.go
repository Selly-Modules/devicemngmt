package devicemngmt

import (
	"context"
	"fmt"

	"github.com/Selly-Modules/logger"
	ua "github.com/mssola/user_agent"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

//  getDeviceCollection ...
func (s Service) getDeviceCollection() *mongo.Collection {
	return s.DB.Collection(fmt.Sprintf("%s-%s", s.TablePrefix, tableDevice))
}

func (s Service) isDeviceIDExisted(ctx context.Context, deviceID string) bool {
	var (
		col    = s.getDeviceCollection()
		device = Device{}
	)

	if err := col.FindOne(ctx, bson.M{"deviceID": deviceID}).Decode(&device); err != nil {
		logger.Error("devicemngt - findByDeviceID", logger.LogData{
			"deviceID": deviceID,
			"err":      err.Error(),
		})
	}
	if !device.ID.IsZero() {
		return true
	}

	return false
}

func getOSName(userAgent string) string {
	uaData := ua.New(userAgent)
	return uaData.OSInfo().Name
}

func getOSVersion(userAgent string) string {
	uaData := ua.New(userAgent)
	return uaData.OSInfo().Version
}

func getLanguage(lang string) string {
	// Language, default is vietnamese(vi)
	if lang == langEn {
		return langEn
	}
	return langVi
}
