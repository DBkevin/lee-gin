package global

import (
	"lee-gin/config"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Application struct {
	ConfigViper *viper.Viper
	Config      config.Configuation
	Log         *zap.Logger
	DB          *gorm.DB
}

var App = new(Application)
