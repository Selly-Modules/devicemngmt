package devicemngmt

import (
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
)

//  getDeviceCollection ...
func (s Service) getDeviceCollection() *mongo.Collection {
	if s.TablePrefix != "" {
		return s.DB.Collection(fmt.Sprintf("%s-%s", s.TablePrefix, tableDevice))
	}
	return s.DB.Collection(tableDevice)
}
