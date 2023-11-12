// Package config Пакет предоставляет функции для работы с конфигурацией сервера и базы данных.
// Использует библиотеку "github.com/spf13/viper" для чтения конфигурационных файлов в формате YAML.
package config

import (
	"github.com/mynamedust/simple-stock/pkg/models"
	"github.com/spf13/viper"
)

const (
	fileName = "config"
	fileType = "yaml"
	filePath = "./internal/config"
)

// New Конструктор конфигурации сервера.
func New() (config models.ServerConfig, err error) {
	viper.SetConfigName(fileName)
	viper.SetConfigType(fileType)
	viper.AddConfigPath(filePath)

	if err = viper.ReadInConfig(); err != nil {
		return
	}
	if err = viper.Unmarshal(&config); err != nil {
		return
	}
	return
}
