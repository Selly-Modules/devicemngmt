package devicemngmt

import (
	ua "github.com/mssola/user_agent"
)

func getUserAgentData(userAgent string) (string, string, bool) {
	uaData := ua.New(userAgent)
	return uaData.OSInfo().Name, uaData.OSInfo().Version, uaData.Mobile()
}

func getLanguage(lang string) string {
	// Language, default is vietnamese(vi)
	if lang == langEn {
		return langEn
	}
	return langVi
}
