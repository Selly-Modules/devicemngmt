package devicemngmt

import (
	"errors"

	"github.com/Selly-Modules/logger"
)

func (co CreateOptions) validate() error {
	// DeviceID
	if co.DeviceID == "" {
		logger.Error("devicemngmt - Create: no deviceID data", logger.LogData{
			"payload": co,
		})
		return errors.New("no deviceID data")
	}

	// UserAgent
	if co.UserAgent == "" {
		logger.Error("devicemngmt - Create: no userAgent data", logger.LogData{
			"payload": co,
		})
		return errors.New("no userAgent data")
	}

	// IP
	if co.IP == "" {
		logger.Error("devicemngmt - Create: no ip data", logger.LogData{
			"payload": co,
		})
		return errors.New("no ip data")
	}

	// UserID
	if co.UserID == "" {
		logger.Error("devicemngmt - Create: no userID data", logger.LogData{
			"payload": co,
		})
		return errors.New("no userID data")
	}

	// AuthToken
	if co.AuthToken == "" {
		logger.Error("devicemngmt - Create: no authToken data", logger.LogData{
			"payload": co,
		})
		return errors.New("no authToken data")
	}

	return nil
}

func (uo UpdateOptions) validate() error {
	// UserAgent
	if uo.UserAgent == "" {
		logger.Error("devicemngmt - Update: no userAgent data", logger.LogData{
			"payload": uo,
		})
		return errors.New("no userAgent data")
	}

	// IP
	if uo.IP == "" {
		logger.Error("devicemngmt - Update: no ip data", logger.LogData{
			"payload": uo,
		})
		return errors.New("no ip data")
	}

	// AuthToken
	if uo.AuthToken == "" {
		logger.Error("devicemngmt - Update: no authToken data", logger.LogData{
			"payload": uo,
		})
		return errors.New("no authToken data")
	}

	return nil
}
