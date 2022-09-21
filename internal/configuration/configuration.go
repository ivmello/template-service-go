package configuration

import (
	"fmt"
	"os"
	"strconv"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/joho/godotenv"
)

type AppEnvEnum string

const (
	Development AppEnvEnum = "DEVELOPMENT"
	Test        AppEnvEnum = "TEST"
	Production  AppEnvEnum = "PRODUCTION"
	Staging     AppEnvEnum = "STAGING"
)

type AppConfig struct {
	Port      string `env:"APP_PORT"`
	Env       string `env:"APP_ENV"`
	IsTesting bool   `env:"APP_TESTING"`
}

func GetEnvString(key string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return ""
}

func GetEnvBool(key string) bool {
	if value, ok := os.LookupEnv(key); ok {
		if value, err := strconv.ParseBool(value); err == nil {
			return value
		}
	}
	return false
}

func (appConfig AppConfig) Validate() error {
	return validation.ValidateStruct(
		&appConfig,
		validation.Field(&appConfig.Port, validation.Required),
		validation.Field(&appConfig.Env, validation.Required),
	)
}

func Load() (*AppConfig, error) {
	config := AppConfig{}
	err := godotenv.Load()
	if err != nil {
		fmt.Println("[WARN] .env file not found. Loading from system environment")
	}
	config.Port = GetEnvString("APP_PORT")
	config.Env = GetEnvString("APP_ENV")
	config.IsTesting = GetEnvBool("APP_TESTING")
	return &config, config.Validate()
}
