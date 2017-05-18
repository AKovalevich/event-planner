package main

import (
	"github.com/AKovalevich/event-planner/models"
	"github.com/AKovalevich/event-planner/apis"
	"github.com/AKovalevich/event-planner/response"
	"github.com/AKovalevich/event-planner/app"

	"log"
	"gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
	"github.com/crgimenes/goConfig"
	"github.com/asaskevich/govalidator"
	_ "github.com/crgimenes/goConfig/toml"
	"fmt"
)

// Initializes all required components and run migrations
func init() {
	// Prepare configuration data
	config := app.Config()
	goConfig.File = "config.toml"
	err := goConfig.Parse(config)
	if err != nil {
		log.Fatal(err)
	}

	// load error messages
	if err := response.LoadMessages("config/errors.yaml"); err != nil {
		panic(fmt.Errorf("Failed to read the error message file: %s", err))
	}

	// Prepare structure validators
	govalidator.SetFieldsRequiredByDefault(true)

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

	app.Post("upload/images", apis.PostImage)
	app.Post("account", apis.PostAccount)
	app.Post("event", apis.PostEvent)
	app.Post("team", apis.PostTeam)
	app.Get("team/:team_id/event", apis.GetTeamEvent)
	app.Get("team/:team_id/account", apis.GetTeamAccount)
	app.Get("team/:team_id", apis.GetTeam)
	app.Post("auth", apis.Auth)

	app.Listen(":8081")
}
