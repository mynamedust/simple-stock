// Package config Пакет предоставляет функции для работы с конфигурацией сервера и базы данных.
// Использует библиотеку "github.com/spf13/viper" для чтения конфигурационных файлов в формате YAML.
package config

import (
	"github.com/spf13/viper"
	"log"
)

// Server Конфигурация сервера, состоящая из прослушиваемого адреса и базы данных.
type Server struct {
	Address  string
	Database Storage
}

// Storage Конфигурация базы данных, включающая в себя данные, требуемые для подключения.
type Storage struct {
	Name     string
	User     string
	Password string
	Host     string
	Port     string
	SSL      string
}

// New Конструктор конфигурации сервера.
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
