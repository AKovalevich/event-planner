package main

import (
	"github.com/AKovalevich/event-planner/models"
	"github.com/AKovalevich/event-planner/handlers"
	"github.com/AKovalevich/event-planner/config"

	"log"
	"gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
	"github.com/crgimenes/goConfig"
	_ "github.com/crgimenes/goConfig/toml"
	"github.com/Sirupsen/logrus"
)

// Initializes all required components and run migrations
func init() {
	// Prepare configuration data
	config := config.Get()
	goConfig.File = "config.toml"
	err := goConfig.Parse(config)
	if err != nil {
		log.Fatal(err)
	}

	log := logrus.New()
	if config.Debug {
		log.Level = logrus.DebugLevel
		log.Debug("Debug mode is on.")
	} else {
		log.Level = logrus.InfoLevel
	}

	// Migrate models
	models.UserMigrate()
	models.AccountMigrate()
	models.ImageMigrate()
	models.TeamMigrate()
	models.EventMigrate()
}

func main() {
	// Serve!
	app := iris.New()
	app.Adapt(httprouter.New())

	app.Post("upload/images", handlers.PostImage);
	app.Post("account", handlers.PostAccount);
	app.Post("event", handlers.PostEvent);
	app.Get("team/:team_id/event", handlers.GetTeamEvent);
	app.Get("team/:team_id/account", handlers.GetTeamAccount);
	app.Get("team/:team_id", handlers.GetTeam);

	app.Listen(":8082")
}
