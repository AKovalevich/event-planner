package apis

import "gopkg.in/kataras/iris.v6"

//
func PostImage(ctx *iris.Context) {
	ctx.JSON(iris.StatusOK, "OK")
}