package global

import (
	"lee-gin/config"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Application struct {
	ConfigViper *viper.Viper
	Config      config.Configuation
	Log         *zap.Logger
}

var App = new(Application)
