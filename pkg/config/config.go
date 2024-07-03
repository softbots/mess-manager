package config

import (
	"log"

	"github.com/spf13/viper"
)

var LocalConfig *Config

type Config struct {
	DBUSER string `mapstructure:"DBUSER"`
	DBPASS string `mapstructure:"DBPASS"`
	DBIP   string `mapstructure:"DBIP"`
	DBNAME string `mapstructure:"DBNAME"`
	PORT   string `mapstructure:"PORT"`
}

func InitConfig() *Config {
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading config", err)
	}

	var config *Config

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal("Error reading env file", err)
	}
	return config
}

func SetConfig() {
	LocalConfig = InitConfig()
}
