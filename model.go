package devicemngmt

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Device ...
type Device struct {
	ID              primitive.ObjectID `bson:"_id" json:"_id"`
	DeviceID        string             `bson:"deviceID" json:"deviceId"` // unique
	IP              string             `bson:"ip" json:"ip"`
	OSName          string             `bson:"osName" json:"osName"`
	OSVersion       string             `bson:"osVersion" json:"osVersion"`
	AppVersion      string             `bson:"appVersion" json:"appVersion"`
	Language        string             `bson:"language" json:"language"` // vi, en
	IsMobile        bool               `bson:"isMobile" json:"isMobile"`
	LastActivatedAt time.Time          `bson:"lastActivatedAt" json:"lastActivatedAt"`
	UserID          primitive.ObjectID `bson:"userID" json:"userId"`
	AuthToken       string             `bson:"authToken" json:"authToken"`
	FCMToken        string             `bson:"fcmToken" json:"fcmToken"`
	CreatedAt       time.Time          `bson:"createdAt" json:"createdAt"`
}
