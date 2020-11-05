package global

import (
	"fsc/config"
	"github.com/spf13/viper"
)

var (
	FSC_CONFIG config.FSCConfig
	FSC_VP     *viper.Viper
)