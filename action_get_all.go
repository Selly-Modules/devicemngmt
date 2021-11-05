package devicemngmt

import (
	"context"

	"github.com/Selly-Modules/mongodb"
	"go.mongodb.org/mongo-driver/bson"
)

// FindAllDevicesByUserID ...
func (s Service) FindAllDevicesByUserID(userID string) []ResponseDevice {
	var (
		ctx    = context.Background()
		col    = s.getDeviceCollection()
		docs   = make([]Device, 0)
		result = make([]ResponseDevice, 0)
		cond   = bson.M{
			"userID": mongodb.NewIDFromString(userID),
		}
	)

	// Find
	cursor, err := col.Find(ctx, cond)
	if err != nil {
		return result
	}
	defer cursor.Close(ctx)
	cursor.All(ctx, &docs)

	// Get response data
	for _, doc := range docs {
		result = append(result, ResponseDevice{
			ID:       doc.ID,
			DeviceID: doc.DeviceID,
			IP:       doc.IP,
			OS: ResponseOS{
				Name:    doc.OSName,
				Version: doc.OSVersion,
			},
			AppVersion:      doc.AppVersion,
			Language:        doc.Language,
			IsMobile:        doc.IsMobile,
			FCMToken:        doc.FCMToken,
			Model:           doc.Model,
			Manufacturer:    doc.Manufacturer,
			LastActivatedAt: doc.LastActivatedAt,
			CreatedAt:       doc.CreatedAt,
		})
	}

	return result
}
