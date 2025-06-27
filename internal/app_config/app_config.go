package appconfig

import (
	"fmt"
	"sync"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

var AppConfigSingleton = AppConfig{}

var once sync.Once

type AppConfig struct {
	DbUser     string `mapstructure:"db_user" validate:"required"`
	DbPassword string `mapstructure:"db_password" validate:"required"`
	DbPort     int    `mapstructure:"db_port" validate:"required"`
	DbHost     string `mapstructure:"db_host" validate:"required"`
	DbName     string `mapstructure:"db_name" validate:"required"`
}

func (a *AppConfig) InitAppConfig() {
	once.Do(func() {
		viper.AutomaticEnv()

		viper.BindEnv("DB_PASSWORD")
		viper.BindEnv("DB_USER")
		viper.BindEnv("DB_PORT")
		viper.BindEnv("DB_HOST")
		viper.BindEnv("DB_NAME")

		err := viper.Unmarshal(a)
		if err != nil {
			panic(fmt.Sprintf("Error found in app config: %v", err))
		}

		validate := validator.New()
		if err := validate.Struct(a); err != nil {
			panic(fmt.Sprintf("Error validating app config: %v", err))
		}
	})
}
