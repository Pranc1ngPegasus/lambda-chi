package configuration

import (
	"io"

	"github.com/kelseyhightower/envconfig"
)

var globalConfig Config

type (
	Config struct {
		ListenPort int `envconfig:"PORT" default:"3000"`
	}
)

func Usage(output io.Writer) {
	if err := envconfig.Usagef("", &globalConfig, output, envconfig.DefaultTableFormat); err != nil {
		panic(err.Error())
	}
}

func Load() {
	envconfig.MustProcess("", &globalConfig)
}

func Get() Config {
	return globalConfig
}
