package config

import (
	"github.com/spf13/viper"
	"log"
)

type Server struct {
	Address  string
	Database Storage
}

type Storage struct {
	Name     string
	User     string
	Password string
	Host     string
	Port     string
	SSL      string
}

func New() (config Server) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatal(err)
	}
	return
}
