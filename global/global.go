package global

import (
	"fsc/config"
	"fsc/model"
	oplogging "github.com/op/go-logging"
	"github.com/spf13/viper"
)

var (
	FSC_CONFIG config.FSCConfig
	FSC_VP     *viper.Viper
	FSC_LOG *oplogging.Logger
	FSC_USER *model.UserStruct
)