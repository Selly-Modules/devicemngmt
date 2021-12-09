package devicemngmt

import (
	"context"
	"errors"

	"github.com/Selly-Modules/logger"
	"github.com/Selly-Modules/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

// FindAllDevicesByUserIDList ...
func (s Service) FindAllDevicesByUserIDList(userIDList []string) (result []Device, err error) {
	var (
		ctx = context.Background()
		col = s.getDeviceCollection()
	)
	result = make([]Device, 0)

	// Validate
	if len(userIDList) <= 0 {
		err = errors.New("user list must be greater than 0")
	}
	if len(userIDList) > limit200 {
		err = errors.New("user list must be less than 200")
	}

	// Get ids
	ids := make([]primitive.ObjectID, 0)
	for _, value := range userIDList {
		if id, valid := mongodb.NewIDFromString(value); valid {
			ids = append(ids, id)
		}
	}

	// Condition
	cond := bson.M{
		"userId": bson.M{
			"$in": ids,
		},
	}

	// Find
	cursor, err := col.Find(ctx, cond)
	if err != nil {
		logger.Error("devicemngmt - findAllDevicesByUserIDList ", logger.LogData{
			"err": err.Error(),
		})
		return
	}
	defer cursor.Close(ctx)
	cursor.All(ctx, &result)

	return result, nil
}
