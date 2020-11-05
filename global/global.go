package global

import (
	"fsc/config"
	"fsc/util"
	"github.com/spf13/viper"
)

type H map[string]interface{}

var (
	FSC_CONFIG config.FSCConfig
	FSC_VP     *viper.Viper
)