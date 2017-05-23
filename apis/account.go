package apis

import (
	"gopkg.in/kataras/iris.v6"
	"github.com/AKovalevich/event-planner/models"
	"github.com/spf13/cast"
	"github.com/AKovalevich/event-planner/app"
	"github.com/AKovalevich/event-planner/response"
	"fmt"
)

//
func PostAccount(ctx *iris.Context) {
	teamID := ctx.Param("team_id")
	// get request scope
	scope := ctx.Get("request_scope")
	user := ctx.Get("User")
	// try to load team
	team, err := models.GetTeam(cast.ToUint(teamID), scope.(app.RequestScope).GetTx())
	if err != nil {
		res := response.InternalServerError("Can't load team", err.Error())
		ctx.JSON(res.StatusCode(), res.Struct())
		return
	}

	if !team.HasUser(cast.ToUint(user.(models.User).ID), scope.(app.RequestScope).GetTx()) {
		res := response.AccessDenied()
		ctx.JSON(res.StatusCode(), res.Struct())
		return
	}

	var account = &models.Account{}
	if err := ctx.ReadJSON(&account); err != nil {
		ctx.JSON(iris.StatusServiceUnavailable, "Service unavailable")
	} else {
		account.Teams = append(account.Teams, *team)

		fmt.Printf("%+v\n", account)
		account, err = models.CreateAccount(account, scope.(app.RequestScope).GetTx())
		if err != nil {
			res := response.InvalidData("")
			ctx.JSON(res.StatusCode(), res.Struct())
			return
		}

		ctx.JSON(iris.StatusOK, account)
	}
}

func GetAccount(ctx *iris.Context) {
	teamID := ctx.Param("team_id")
	// get request scope
	scope := ctx.Get("request_scope")
	user := ctx.Get("User")
	// try to load team
	team, err := models.GetTeam(cast.ToUint(teamID), scope.(app.RequestScope).GetTx())
	if err != nil {
		res := response.InternalServerError("Can't load team", err.Error())
		ctx.JSON(res.StatusCode(), res.Struct())
		return
	}

	if !team.HasUser(cast.ToUint(user.(models.User).ID), scope.(app.RequestScope).GetTx()) {
		res := response.AccessDenied()
		ctx.JSON(res.StatusCode(), res.Struct())
		return
	}

	//@TODO Use paginated list
}