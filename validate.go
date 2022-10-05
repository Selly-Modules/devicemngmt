package devicemngmt

import (
	"errors"

	"github.com/Selly-Modules/logger"
	"github.com/Selly-Modules/mongodb"
)

func (co CreateOptions) validate() error {
	// DeviceID
	if co.DeviceID == "" {
		logger.Error("devicemngmt - Create: no deviceID data", logger.LogData{
			Source:  "devicemngmt.validate",
			Message: "devicemngmt - Create: no deviceID data",
			Data:    co,
		})
		return errors.New("no deviceID data")
	}

	// IP
	if co.IP == "" {
		logger.Error("devicemngmt - Create: no ip data", logger.LogData{
			Source:  "devicemngmt.validate",
			Message: "devicemngmt - Create: no ip data",
			Data:    co,
		})
		return errors.New("no ip data")
	}

	// UserID
	if co.UserID == "" {
		logger.Error("devicemngmt - Create: no userID data", logger.LogData{
			Source:  "devicemngmt.validate",
			Message: "devicemngmt - Create: no userID data",
			Data:    co,
		})
		return errors.New("no userID data")
	}
	if _, isValid := mongodb.NewIDFromString(co.UserID); !isValid {
		return errors.New("invalid userID data")
	}

	// AuthToken
	if co.AuthToken == "" {
		logger.Error("devicemngmt - Create: no authToken data", logger.LogData{
			Source:  "devicemngmt.validate",
			Message: "devicemngmt - Create: no authToken data",
			Data:    co,
		})
		return errors.New("no authToken data")
	}

	return nil
}

func (uo UpdateOptions) validate() error {
	// UserID
	if uo.UserID == "" {
		logger.Error("devicemngmt - Update: no userID data", logger.LogData{
			Source:  "devicemngmt.validate",
			Message: "devicemngmt - Update: no userID data",
			Data:    uo,
		})
		return errors.New("no userID data")
	}
	if _, isValid := mongodb.NewIDFromString(uo.UserID); !isValid {
		return errors.New("invalid userID data")
	}

	// IP
	if uo.IP == "" {
		logger.Error("devicemngmt - Update: no ip data", logger.LogData{
			Source:  "devicemngmt.validate",
			Message: "devicemngmt - Update: no ip data",
			Data:    uo,
		})
		return errors.New("no ip data")
	}

	// AuthToken
	if uo.AuthToken == "" {
		logger.Error("devicemngmt - Update: no authToken data", logger.LogData{
			Source:  "devicemngmt.validate",
			Message: "devicemngmt - Update: no authToken data",
			Data:    uo,
		})
		return errors.New("no authToken data")
	}

	return nil
}
