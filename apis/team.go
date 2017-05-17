package apis

import (
	"github.com/AKovalevich/event-planner/models"

	"gopkg.in/kataras/iris.v6"
	"github.com/AKovalevich/event-planner/response"
)

// Create new team
// Method: POST
func PostTeam(ctx *iris.Context) {
	team := &models.Team{}

	if err := ctx.ReadJSON(&team); err != nil {
		ctx.JSON(iris.StatusServiceUnavailable, "Service unavailable")
	} else {
		team, err := models.CreateTeam(team)
		if err != nil {
			res := response.InternalServerError("Can't create team", err.Error())
			ctx.JSON(res.StatusCode(), res.Struct())

		} else {
			ctx.JSON(iris.StatusOK, team)
		}
	}
}

//
func GetTeamEvent(ctx *iris.Context) {
	ctx.JSON(iris.StatusOK, "OK")
}

//
func GetTeamAccount(ctx *iris.Context) {
	team_id := ctx.Param("team_id")
	team, err := models.GetTeam(team_id)
	if err != nil {
		res := response.InternalServerError("Can't load team", err.Error())
		ctx.JSON(res.StatusCode(), res.Struct())
		return
	}

	ctx.JSON(iris.StatusOK, team)
}

//
func GetTeam(ctx *iris.Context) {
	team_id := ctx.Param("team_id")
	team, err := models.GetTeam(team_id)
	if err != nil {
		res := response.InternalServerError("Can't load team", err.Error())
		ctx.JSON(res.StatusCode(), res.Struct())
		return
	}

	ctx.JSON(iris.StatusOK, team)
}