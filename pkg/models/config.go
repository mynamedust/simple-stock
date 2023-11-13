package models

// ServerConfig конфигурация сервера.
type ServerConfig struct {
	Address  string
	Database StorageConfig
}

// StorageConfig конфигурация базы данных.
type StorageConfig struct {
	Name     string
	User     string
	Password string
	Host     string
	Port     string
	SSL      string
}
