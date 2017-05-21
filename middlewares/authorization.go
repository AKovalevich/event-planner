package middlewares

import (
	"github.com/dgrijalva/jwt-go"
	jwtmiddleware "github.com/iris-contrib/middleware/jwt"
	"github.com/AKovalevich/event-planner/app"
	"gopkg.in/kataras/iris.v6"
	"github.com/AKovalevich/event-planner/response"
	"strings"
	"fmt"
	"github.com/AKovalevich/event-planner/models"
)

type authorizationMiddleware struct {}

// serve serves the middleware
func NewAuthorization() *jwtmiddleware.Middleware {
	return jwtmiddleware.New(jwtmiddleware.Config{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(app.Config().Secret), nil
		},
		// When set, the middleware verifies that tokens are signed with the specific signing algorithm
		// If the signing method is not constant the ValidationKeyGetter callback can be used to implement additional checks
		// Important to avoid security issues described here: https://auth0.com/blog/2015/03/31/critical-vulnerabilities-in-json-web-token-libraries/
		SigningMethod: jwt.SigningMethodHS256,
		ErrorHandler: errorHandler,
		Extractor: FromAuthHeader,
	})
}

// fromAuthHeader is a "TokenExtractor" that takes a give context and extracts
// the JWT token from the Authorization header.
func FromAuthHeader(ctx *iris.Context) (string, error) {

	authHeader := ctx.RequestHeader("Authorization")
	if authHeader == "" {
		return "", nil // No error, just no token
	}

	// TODO: Make this a bit more robust, parsing-wise
	authHeaderParts := strings.Split(authHeader, " ")
	if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
		return "", fmt.Errorf("Authorization header format must be Bearer {token}")
	}

	// validate that it valid token
	token, err := models.LoadToken(authHeaderParts[1])

	if err != nil {
		return "", fmt.Errorf("Service unavailable")
	}

	if token.Token == "" {
		return "", fmt.Errorf("Invalid token")
	}

	if token.Expired() {
		return "", fmt.Errorf("The token has expired")
	}

	if err := token.User.LoadUserAssociations(); err != nil {
		return "", fmt.Errorf("Service unavailable")
	}

	ctx.Set("User", token.User)

	return token.Token, nil
}

func errorHandler(ctx *iris.Context, string string) {
	res := response.Unauthorized(string)
	ctx.JSON(res.StatusCode(), res.Struct())
	return
}
