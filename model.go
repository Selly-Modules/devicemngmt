package devicemngmt

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Device ...
type Device struct {
	ID             primitive.ObjectID `bson:"_id"`
	DeviceID       string             `bson:"deviceID"` // unique
	IP             string             `bson:"ip"`
	OSName         string             `bson:"osName"`
	OSVersion      string             `bson:"osVersion"`
	AppVersion     string             `bson:"appVersion"`
	Language       string             `bson:"language"` // vi, en
	IsMobile       bool               `bson:"isMobile"`
	LastActivityAt time.Time          `bson:"lastActivityAt"`
	UserID         primitive.ObjectID `bson:"userID"`
	AuthToken      string             `bson:"authToken"`
	FCMToken       string             `bson:"fcmToken"`
	CreatedAt      time.Time          `bson:"createdAt"`
}
