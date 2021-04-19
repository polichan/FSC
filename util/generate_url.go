package util

import "fsc/global"

func GenerateUrl(api string) string {
	return global.FSC_CONFIG.SportCampusConfig.Api + api
}
