package config

import (
	validator "github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Validator interface {
	Struct(interface{}) error
}

type Configuration struct {
	validator Validator `validate:"required"`

	App *App `validate:"required"`
}

func (c *Configuration) Validate() (err error) {
	err = c.validator.Struct(c)

	return
}

func New() (c *Configuration) {
	viper.SetConfigFile(".env")

	viper.SetDefault("WINMETRICS_LISTEN_ADDRESS", DefaultListenAddress)

	err := viper.ReadInConfig()
	if err != nil {
		logrus.Warnf("Configuration error while reading .env file: %v, trying to read from environment variable instead", err)

		handleError(viper.BindEnv("WINMETRICS_LISTEN_ADDRESS"))
	}

	c = &Configuration{}
	c.validator = validator.New()

	c.App = &App{}
	c.App.ListenAddress = viper.GetString("WINMETRICS_LISTEN_ADDRESS")

	return
}

func handleError(err error) {
	if err != nil {
		logrus.Fatalf("Configuration error: %v", err)
	}
}
