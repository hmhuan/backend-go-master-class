package util

import "github.com/spf13/viper"

type Config struct {
	DbDriver      string `mapstructure:"DB_DRIVER"`
	DbSource      string `mapstructure:"DB_SOURCE"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
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
