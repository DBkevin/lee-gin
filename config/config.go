package config

type Configuation struct {
	App App `mapstructure:"app" json:"app" yaml:"app"`
}
