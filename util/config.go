package util

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DbDriver            string        `mapstructure:"DB_DRIVER"`
	DbSource            string        `mapstructure:"DB_SOURCE"`
	ServerAddress       string        `mapstructure:"SERVER_ADDRESS"`
	TokenSimmetricKey   string        `mapstructure:"TOKEN_SIMMETRIC_KEY"`
	AccessTokenDuration time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
}

func LoadConfig(path string) (*Config, error) {

	var config Config
	// config path and file for viper
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env") // yml, json, toml, properties ...

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&config)
	return &config, err
}
