package middlewares

import (
	"github.com/AKovalevich/event-planner/utils"
	"github.com/AKovalevich/event-planner/app"

	"gopkg.in/kataras/iris.v6"
	"github.com/jinzhu/gorm"
	"time"
	"log"
)

type requestScopeMiddleware struct {}

// Serve serves the middleware
func (l *requestScopeMiddleware) Serve(ctx *iris.Context) {
	// prepare request scope
	scope := app.NewRequestScope(time.Now(), ctx)

	// create new db connection
	db, err := utils.GetDB()
	if err != nil {
		log.Fatal(err)
	}
	tx := db.Begin()
	scope.SetTx(tx)

	ctx.Set("request_scope", scope)
	ctx.Next()

	if scope.GetRollback() {
		scope.GetTx().(*gorm.DB).Rollback()
	} else {
		scope.GetTx().(*gorm.DB).Commit()
	}
}

// New returns the logger middleware
// receives optional configs(logger.Config)
func NewRequestScope() iris.HandlerFunc {
	l := &requestScopeMiddleware{}

	return l.Serve
}