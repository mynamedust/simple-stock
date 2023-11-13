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
func New() (models.ServerConfig, error) {
	viper.SetConfigName(fileName)
	viper.SetConfigType(fileType)
	viper.AddConfigPath(filePath)

	var config models.ServerConfig
	if err := viper.ReadInConfig(); err != nil {
		return config, err
	}
	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}
	return config, nil
}
