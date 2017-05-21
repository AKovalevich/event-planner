package app

import "sync"

// Application configuration structure, contains main parameters for application
// Database configuration
type dataBase struct {
	Host string `json:"host" toml:"host" cfg:"host" cfgDefault:"localhost" env:"SERVICE_STORAGE_HOST"`
	Port int    `json:"port" toml:"port" cfg:"port" cfgDefault:"3306" env:"SERVICE_STORAGE_PORT"`
	User string `json:"user" toml:"user" cfg:"user" cfgDefault:"root" env:"SERVICE_STORAGE_USER"`
	Password string `json:"password" toml:"password" cfg:"password" cfgDefault:"" env:"SERVICE_STORAGE_PASSWORD"`
	Name string `json:"name" toml:"name" cfg:"name" cfgDefault:"" env:"SERVICE_STORAGE_NAME"`
}

// Main Application configuration
type appConfig struct {
	Port string `json:"port" toml:"port" cfg:"port" cfgDefault:"8081" env:"SERVICE_PORT"`
	Domain string `json:"domain" toml:"domain" cfg:"domain" cfgDefault:"localhost" env:"SERVICE_DOMAIN"`
	Debug bool `json:"debug" toml:"debug" cfg:"debug" cfgDefault:"false" env:"SERVICE_DEBUG"`
	DataBase dataBase `json:"database" toml:"database" cfg:"database" env:"SERVICE_DATABASE"`
	Secret string `json:"secret" toml:"secret" cfg:"secret" env:"SERVICE_SERCRET_KEY"`
}

var config *appConfig
var once sync.Once

//
func Config() *appConfig {
	once.Do(func() {
		config = &appConfig{}
	})
	return config
}
