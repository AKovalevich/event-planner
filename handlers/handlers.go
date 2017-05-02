package handlers

import (
	"github.com/AKovalevich/event-planner/models"

	"gopkg.in/kataras/iris.v6"
)

func PostImage(ctx *iris.Context) {
	ctx.JSON(iris.StatusOK, "OK")
}
func PostEvent(ctx *iris.Context) {
	ctx.JSON(iris.StatusOK, "OK")
}
func GetTeamEvent(ctx *iris.Context) {
	ctx.JSON(iris.StatusOK, "OK")
}
func GetTeamAccount(ctx *iris.Context) {
	ctx.JSON(iris.StatusOK, "OK")
}
func GetTeam(ctx *iris.Context) {
	ctx.JSON(iris.StatusOK, "OK")
}
func PostAccount(ctx *iris.Context) {
	ctx.JSON(iris.StatusOK, "OK")
}
func PostTeam(ctx *iris.Context) {
	team := &models.Team{}

	if err := ctx.ReadJSON(&team); err != nil {
		ctx.JSON(iris.StatusBadRequest, err.Error())
	} else {
		team, err := models.CreateTeam(team)

		if err != nil {
			ctx.JSON(iris.StatusServiceUnavailable, err.Error())
		}

		ctx.JSON(iris.StatusOK, team)
	}
}
