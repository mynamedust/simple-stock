package models

// ServerConfig Конфигурация сервера, состоящая из прослушиваемого адреса и базы данных.
type ServerConfig struct {
	Address  string
	Database StorageConfig
}

// StorageConfig Конфигурация базы данных, включающая в себя данные, требуемые для подключения.
type StorageConfig struct {
	Name     string
	User     string
	Password string
	Host     string
	Port     string
	SSL      string
}
