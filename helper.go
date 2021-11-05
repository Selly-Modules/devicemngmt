package devicemngmt

import (
	"fmt"

	ua "github.com/mssola/user_agent"
	"go.mongodb.org/mongo-driver/mongo"
)

//  getDeviceCollection ...
func (s Service) getDeviceCollection() *mongo.Collection {
	return s.DB.Collection(fmt.Sprintf("%s-%s", s.TablePrefix, tableDevice))
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
