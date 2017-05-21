package app

import (
	"gopkg.in/kataras/iris.v6"
	"time"
)

// RequestScope contains the application-specific information that are carried around in a request.
type RequestScope interface {
	// UserID returns the ID of the user for the current request
	GetUser() interface{}
	// SetUserID sets the ID of the currently authenticated user
	SetUser(user interface{})
	// RequestID returns the ID of the current request
	GetRequestID() string
	// Tx returns the currently active database transaction that can be used for DB query purpose
	GetTx() interface{}
	// SetTx sets the database transaction
	SetTx(tx interface{})
	// Rollback returns a value indicating whether the current database transaction should be rolled back
	GetRollback() bool
	// SetRollback sets a value indicating whether the current database transaction should be rolled back
	SetRollback(v bool)
	// Now returns the timestamp representing the time when the request is being processed
	Now() time.Time
}

type requestScope struct {
	now time.Time // the time when the request is being processed
	requestID string    // an ID identifying one or multiple correlated HTTP requests
	user interface{}    // an ID identifying the current user
	rollback bool      // whether to roll back the current transaction
	tx interface{}   // the currently active transaction
}

func (rs *requestScope) GetUser() interface{} {
	return rs.user
}

func (rs *requestScope) SetUser(user interface{}) {
	rs.user = user
}

func (rs *requestScope) GetRequestID() string {
	return rs.requestID
}

func (rs *requestScope) GetTx() interface{} {
	return rs.tx
}

func (rs *requestScope) SetTx(tx interface{}) {
	rs.tx = tx
}

func (rs *requestScope) GetRollback() bool {
	return rs.rollback
}

func (rs *requestScope) SetRollback(v bool) {
	rs.rollback = v
}

func (rs *requestScope) Now() time.Time {
	return rs.now
}

// newRequestScope creates a new RequestScope with the current request information.
func NewRequestScope(now time.Time, ctx *iris.Context) RequestScope {
	requestID := ctx.Header().Get("X-Request-Id")
	return &requestScope{
		now:       now,
		requestID: requestID,
	}
}
