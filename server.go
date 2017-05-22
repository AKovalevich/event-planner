package main

import (
	"github.com/AKovalevich/event-planner/middlewares"
	"github.com/AKovalevich/event-planner/response"
	"github.com/AKovalevich/event-planner/models"
	"github.com/AKovalevich/event-planner/apis"
	"github.com/AKovalevich/event-planner/app"

	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
	"gopkg.in/kataras/iris.v6/middleware/logger"
	"github.com/asaskevich/govalidator"
	"github.com/jinzhu/configor"
	"gopkg.in/kataras/iris.v6"
	"log"
	"fmt"
)

// Initializes all required components and run migrations
func init() {

	// Prepare Application configuration
	println("Parse configuration file - config.toml...")
	config := app.Config()
	err := configor.Load(config, "config.toml")
	if err != nil {
		log.Fatal(err)
	}
	println("Done")

	println("Prepare srror messages...")
	// load error messages
	if err := response.LoadMessages("config/errors.yaml"); err != nil {
		panic(fmt.Errorf("Failed to read the error message file: %s", err))
	}
	println("Done")

	// prepare structure validators
	govalidator.SetFieldsRequiredByDefault(true)

	println("Start model migrations...")
	// migrate models
	models.UserMigrate()
	models.AccountMigrate()
	models.ImageMigrate()
	models.TeamMigrate()
	models.EventMigrate()
	models.TokenMigrate()
	println("Done")
}

func main() {
	// serve!
	server := iris.New()
	server.Adapt(httprouter.New())

	if app.Config().Debug {
		server.Adapt(iris.DevLogger())
		customLogger := logger.New(logger.Config{
			// Status displays status code
			Status: true,
			// IP displays request's remote address
			IP: true,
			// Method displays the http method
			Method: true,
			// Path displays the request path
			Path: true,
		})

		server.Use(customLogger)
	}

	scopeHandler := middlewares.NewRequestScope()
	authEndpoints := server.Party("/api/v1/auth/")
	authEndpoints.Use(scopeHandler)
	{
		authEndpoints.Post("token", apis.AuthToken)
		authEndpoints.Post("token/refresh", apis.AuthTokenRefresh)
		authEndpoints.Post("register", apis.AuthRegister)
	}

	authHandler := middlewares.NewAuthorization()
	authorizedEndpoints := server.Party("/api/v1/")
	authorizedEndpoints.Use(scopeHandler)
	authorizedEndpoints.Use(authHandler)
	{
		authorizedEndpoints.Post("account", apis.PostAccount)
		authorizedEndpoints.Post("event", apis.PostEvent)
		authorizedEndpoints.Post("team", apis.PostTeam)
		authorizedEndpoints.Post("team/:team_id/image", apis.PostImage)
		authorizedEndpoints.Get("team/:team_id/event", apis.GetTeamEvent)
		authorizedEndpoints.Get("team/:team_id/account", apis.GetTeamAccount)
		authorizedEndpoints.Get("team/:team_id", apis.GetTeam)
		authorizedEndpoints.Get("user/:user_id/team", apis.GetUserTeam)
	}

	server.Listen(":8081")
}
