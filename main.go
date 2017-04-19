package main

import (
	"github.com/AKovalevich/event-planner/models"

	"gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
	"github.com/crgimenes/goConfig"
	_ "github.com/crgimenes/goConfig/toml"
)

// Database configuration
type dataBase struct {
	Host string `json:"host" toml:"host" cfg:"host" cfgDefault:"localhost"`
	Port int    `json:"port" toml:"port" cfg:"port" cfgDefault:"3306"`
	User string `json:"user" toml:"user" cfg:"user" cfgDefault:"root"`
	Password string `json:"password" toml:"password" cfg:"password" cfgDefault:""`
}

// Main Application configuration
type config struct {
	Domain string `json:"domain" toml:"domain" cfg:"domain" cfgDefault:"localhost"`
	DebugMode bool `json:"debug" toml:"debug" cfg:"debug" cfgDefault:"false"`
	DataBase dataBase `json:"database" toml:"database" cfg:"database"`
}

func main() {
	// Prepare configuration data
	config := config{}
	goConfig.File = "config.toml"
	err := goConfig.Parse(&config)
	if err != nil {
		panic(err.Error())
	}

	models.ImageMigrate()
	models.TeamMigrate()
	models.EventMigrate()

	// Serve!
	app := iris.New()
	app.Adapt(httprouter.New())

	app.HandleFunc("GET", "/", func(ctx *iris.Context) {
		ctx.Writef("OK")
	})

	//app.Listen(":8080")
}
