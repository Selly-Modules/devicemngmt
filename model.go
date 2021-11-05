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
	Model           string             `bson:"model,omitempty" json:"model,omitempty"`
	Manufacturer    string             `bson:"manufacturer,omitempty" json:"manufacturer,omitempty"`
	CreatedAt       time.Time          `bson:"createdAt" json:"createdAt"`
}

// ResponseOS ...
type ResponseOS struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

// ResponseDevice ...
type ResponseDevice struct {
	ID              primitive.ObjectID `json:"_id"`
	DeviceID        string             `json:"deviceId"`
	IP              string             `json:"ip"`
	OS              ResponseOS         `json:"os"`
	AppVersion      string             `json:"appVersion"`
	Language        string             `json:"language"`
	IsMobile        bool               `json:"isMobile"`
	FCMToken        string             `json:"fcmToken"`
	Model           string             `json:"model,omitempty"`
	Manufacturer    string             `json:"manufacturer,omitempty"`
	LastActivatedAt time.Time          `json:"lastActivatedAt"`
	CreatedAt       time.Time          `json:"createdAt"`
}
