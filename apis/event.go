package apis

import (
	"github.com/AKovalevich/event-planner/models"
	"github.com/AKovalevich/event-planner/app"
	"github.com/AKovalevich/event-planner/response"

	"gopkg.in/kataras/iris.v6"
	"github.com/spf13/cast"
)

//
func PostEvent(ctx *iris.Context) {
	// get params
	teamID := ctx.Param("team_id")
	user := ctx.Get("User")
	accountID := ctx.Param("account_id")

	// get request scope
	scope := ctx.Get("request_scope")

	team, err := models.GetTeam(cast.ToUint(teamID), scope.(app.RequestScope).GetTx())
	if err != nil {
		res := response.NotFound("Team")
		ctx.JSON(res.StatusCode(), res.Struct())
		return
	}

	if team.HasUser(cast.ToUint(user.(models.User).ID), scope.(app.RequestScope).GetTx()) {
		// Check account ID
		if team.HasAccount(cast.ToUint(accountID), scope.(app.RequestScope).GetTx()) {
			event := &models.Event{}
			if err := ctx.ReadJSON(&event); err != nil {
				ctx.JSON(iris.StatusServiceUnavailable, "Validation error")
			} else {
				event.TeamID = cast.ToUint(teamID)
				event.UserID = cast.ToUint(cast.ToUint(accountID))
				event.AccountId = cast.ToUint(accountID)
				event, err := models.CreateEvent(event, scope.(app.RequestScope).GetTx())
				if err != nil {
					res := response.InternalServerError("Can't create event", err.Error())
					ctx.JSON(res.StatusCode(), res.Struct())
				} else {
					ctx.JSON(iris.StatusOK, event)
				}
			}
		} else {
			res := response.AccessDenied()
			ctx.JSON(res.StatusCode(), res.Struct())
		}
	} else {
		res := response.AccessDenied()
		ctx.JSON(res.StatusCode(), res.Struct())
	}

}
