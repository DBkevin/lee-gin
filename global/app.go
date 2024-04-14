package global

import (
	"lee-gin/config"

	"github.com/spf13/viper"
)

type Application struct {
	ConfigViper *viper.Viper
	Config      config.Configuation
}

var App=new(Application)