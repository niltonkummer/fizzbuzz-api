package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	HTTPServerHost   string `mapstructure:"HTTP_SERVER_HOST"`
	RedisAddress     string `mapstructure:"REDIS_ADDRESS"`
	RedisPassword    string `mapstructure:"REDIS_PASSWORD"`
	StorageType      string `mapstructure:"STORAGE_TYPE"`
	UseFizzbuzzCache bool   `mapstructure:"USE_FIZZBUZZ_CACHE"`
}

func LoadConfig(path string) Config {
	
	var config Config
	viper.AddConfigPath(path)
	viper.SetConfigName(fmt.Sprintf("server.%s.env", os.Getenv("ENV")))
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic("Error reading config file: " + err.Error())
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		panic("Error unmarshalling config: " + err.Error())
	}

	return config
}
