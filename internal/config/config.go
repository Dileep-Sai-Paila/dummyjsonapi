package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DummyJsonURL string `mapstructure:"DUMMY_JSON_URL"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file, %s", err)
		return nil, err
	}

	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
		return nil, err
	}

	return &config, nil
}
