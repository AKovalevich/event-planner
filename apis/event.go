package apis

import "gopkg.in/kataras/iris.v6"

//
func PostEvent(ctx *iris.Context) {
	ctx.JSON(iris.StatusOK, "OK")
}
