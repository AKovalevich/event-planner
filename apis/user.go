package apis

import (
	"github.com/AKovalevich/event-planner/response"
	"github.com/AKovalevich/event-planner/models"

	"gopkg.in/kataras/iris.v6"
)

//
func GetUserTeam(ctx *iris.Context) {
	user := ctx.Get("User")
	userId, err := ctx.ParamInt("user_id")
	if err != nil {
		res := response.InternalServerError("", err.Error())
		ctx.JSON(res.StatusCode(), res.Struct())
		return
	}

	if int(user.(models.User).ID) != userId {
		res := response.AccessDenied()
		ctx.JSON(res.StatusCode(), res.Struct())
		return
	}

	ctx.JSON(iris.StatusOK, user.(models.User).Teams)
}
