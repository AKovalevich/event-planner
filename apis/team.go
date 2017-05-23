package apis

import (
	"github.com/AKovalevich/event-planner/response"
	"github.com/AKovalevich/event-planner/models"
	"github.com/AKovalevich/event-planner/app"

	"gopkg.in/kataras/iris.v6"
	"github.com/spf13/cast"
)

// Create new team
// Method: POST
func PostTeam(ctx *iris.Context) {
	// get request scope
	scope := ctx.Get("request_scope")

	team := &models.Team{}

	if err := ctx.ReadJSON(&team); err != nil {
		ctx.JSON(iris.StatusServiceUnavailable, "Service unavailable")
	} else {
		team, err := models.CreateTeam(team, scope.(app.RequestScope).GetTx())
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
	// get request scope
	scope := ctx.Get("request_scope")

	team_id := ctx.Param("team_id")
	team, err := models.GetTeam(cast.ToUint(team_id), scope.(app.RequestScope).GetTx())
	if err != nil {
		res := response.InternalServerError("Can't load team", err.Error())
		ctx.JSON(res.StatusCode(), res.Struct())
		return
	}

	ctx.JSON(iris.StatusOK, team)
}

//
func GetTeam(ctx *iris.Context) {
	user := ctx.Get("User")
	// get request scope
	scope := ctx.Get("request_scope")

	team_id := ctx.Param("team_id")
	team, err := models.GetTeam(cast.ToUint(team_id), scope.(app.RequestScope).GetTx())
	if err != nil {
		res := response.InternalServerError("Can't load team", err.Error())
		ctx.JSON(res.StatusCode(), res.Struct())
		return
	}

	if team.HasUser(user.(models.User).ID, scope.(app.RequestScope).GetTx()) {
		ctx.JSON(iris.StatusOK, team)
	} else {
		res := response.AccessDenied()
		ctx.JSON(res.StatusCode(), res.Struct())
	}
}
