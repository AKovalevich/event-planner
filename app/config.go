package app

import "sync"

// Application configuration structure, contains main parameters for application
// Database configuration
type dataBase struct {
	Host string `json:"host" toml:"host" cfg:"host" cfgDefault:"localhost"`
	Port int    `json:"port" toml:"port" cfg:"port" cfgDefault:"3306"`
	User string `json:"user" toml:"user" cfg:"user" cfgDefault:"root"`
	Password string `json:"password" toml:"password" cfg:"password" cfgDefault:""`
	Name string `json:"name" toml:"name" cfg:"name" cfgDefault:""`
}

// Main Application configuration
type appConfig struct {
	Domain string `json:"domain" toml:"domain" cfg:"domain" cfgDefault:"localhost"`
	Debug bool `json:"debug" toml:"debug" cfg:"debug" cfgDefault:"false"`
	DataBase dataBase `json:"database" toml:"database" cfg:"database"`
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
