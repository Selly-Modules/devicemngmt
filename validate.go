package devicemngmt

import (
	"errors"

	"github.com/Selly-Modules/logger"
)

func (co CreateOptions) validate() error{
	// DeviceID
	if co.DeviceID == "" {
		logger.Error("devicemngt - Create: no deviceID data", logger.LogData{
			"payload": co,
		})
		return errors.New("no deviceID data")
	}

	// UserAgent
	if co.UserAgent == "" {
		logger.Error("devicemngt - Create: no userAgent data", logger.LogData{
			"payload": co,
		})
		return  errors.New("no userAgent data")
	}

	// IP
	if co.IP == "" {
		logger.Error("devicemngt - Create: no ip data", logger.LogData{
			"payload": co,
		})
		return errors.New("no ip data")
	}

	// UserID
	if co.UserID == "" {
		logger.Error("devicemngt - Create: no userID data", logger.LogData{
			"payload": co,
		})
		return errors.New("no userID data")
	}

	// AuthToken
	if co.AuthToken == "" {
		logger.Error("devicemngt - Create: no authToken data", logger.LogData{
			"payload": co,
		})
		return errors.New("no authToken data")
	}

	return nil
}
