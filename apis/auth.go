package apis

import (
	"github.com/AKovalevich/event-planner/response"
	"github.com/AKovalevich/event-planner/models"
	"github.com/AKovalevich/event-planner/utils"

	"gopkg.in/kataras/iris.v6"
	"time"
	"fmt"
	"github.com/AKovalevich/event-planner/app"
)

//
func AuthRegister(ctx *iris.Context) {
	scope := ctx.Get("request_scope")
	credentials := models.Credentials{}
	err := ctx.ReadJSON(&credentials)
	if err != nil {
		res := response.InternalServerError("", err.Error())
		ctx.JSON(res.StatusCode(), res.Struct())
		return
	}

	existingUser, err := models.GetUserByEmail(credentials.Email, scope.(app.RequestScope).GetTx())
	if err != nil {
		res := response.InternalServerError("", err.Error())
		ctx.JSON(res.StatusCode(), res.Struct())
		return
	}

	if existingUser.Email != "" {
		res := response.EntityAlreadyExists("User")
		ctx.JSON(res.StatusCode(), res.Struct())
		return
	}

	// create new user via credentials
	user := &models.User{
		Email: credentials.Email,
		Password: utils.HashMd5(credentials.Password),
	}
	user, err = models.CreateUser(user, scope.(app.RequestScope).GetTx())
	if err != nil {
		res := response.InternalServerError("", err.Error())
		ctx.JSON(res.StatusCode(), res.Struct())
		return
	}

	// create new team
	var users = []models.User{}
	users = append(users, *user)
	var team = &models.Team{
		Name: fmt.Sprintf("team-%d", user.ID),
		Status: false,
		Users: users,
	}
	team, err = models.CreateTeam(team, scope.(app.RequestScope).GetTx())
	if err != nil {
		res := response.InternalServerError("", err.Error())
		ctx.JSON(res.StatusCode(), res.Struct())
		return
	}

	// generate token for new user
	expireToken := time.Now().Add(time.Hour * 666).Unix()
	signedToken, err := models.GenerateToken(user, ctx.Host(), expireToken)
	if err != nil {
		res := response.InternalServerError("", err.Error())
		ctx.JSON(res.StatusCode(), res.Struct())
		return
	}

	token, err := models.SaveToken(signedToken, expireToken, user, scope.(app.RequestScope).GetTx())
	if err != nil {
		res := response.InternalServerError("", err.Error())
		ctx.JSON(res.StatusCode(), res.Struct())
		return
	}

	ctx.JSON(iris.StatusOK, iris.Map{"token": signedToken, "refresh_token": token.RefreshToken, "expires at": token.ExpiresAt})
}

//
func AuthTokenRefresh(ctx *iris.Context) {
	scope := ctx.Get("request_scope")
	var tokenRefresh = struct{
		Token string `json:"token"`
		RefreshToken string `json:"refresh_token"`
	}{}

	err := ctx.ReadJSON(&tokenRefresh)
	if err != nil {
		res := response.InternalServerError("", err.Error())
		ctx.JSON(res.StatusCode(), res.Struct())
		return
	}

	token, err := models.LoadToken(tokenRefresh.Token, scope.(app.RequestScope).GetTx())
	if err != nil {
		res := response.InternalServerError("", err.Error())
		ctx.JSON(res.StatusCode(), res.Struct())
		return
	}

	if token.Token == "" || token.RefreshToken == "" {
		res := response.NotFound("Token")
		ctx.JSON(res.StatusCode(), res.Struct())
		return
	}

	if token.RefreshToken != tokenRefresh.RefreshToken {
		res := response.Unauthorized("Invalid token")
		ctx.JSON(res.StatusCode(), res.Struct())
		return
	}

	// refresh token
	if err != nil {
		res := response.InternalServerError("", err.Error())
		ctx.JSON(res.StatusCode(), res.Struct())
		return
	}

	// update token
	err = token.UpdateToken(&token.User, ctx.Host(), scope.(app.RequestScope).GetTx())
	if err != nil {
		res := response.InternalServerError("", err.Error())
		ctx.JSON(res.StatusCode(), res.Struct())
		return
	}

	ctx.JSON(iris.StatusOK, iris.Map{"token": token.Token, "refresh_token": token.RefreshToken, "expires at": token.ExpiresAt})
	return
}

//
func AuthToken(ctx *iris.Context) {
	scope := ctx.Get("request_scope")
	credentials := models.Credentials{}
	err := ctx.ReadJSON(&credentials)
	if err != nil {
		res := response.InternalServerError("", err.Error())
		ctx.JSON(res.StatusCode(), res.Struct())
		return
	}

	// prepare password hash
	hashedPassword := utils.HashMd5(credentials.Password)

	// try to load user by password and email
	user, err := models.GetUserByEmail(credentials.Email, scope.(app.RequestScope).GetTx())
	if err != nil {
		res := response.NotFound("User")
		ctx.JSON(res.StatusCode(), res.Struct())
		return
	}

	if credentials.Email == user.Email && utils.HashMd5(credentials.Password) == hashedPassword { // "test" - "098f6bcd4621d373cade4e832627b4f6"
		// in case if already exists return token
		existingToken, err := models.LoadTokenQuery(struct{id uint}{id: user.TokenID}, scope.(app.RequestScope).GetTx())
		if err != nil {
			res := response.InternalServerError("", err.Error())
			ctx.JSON(res.StatusCode(), res.Struct())
			return
		}

		var token = &models.Token{}

		if existingToken.Token != "" {
			token.Token = existingToken.Token
			token.RefreshToken = existingToken.RefreshToken
			token.ExpiresAt = existingToken.ExpiresAt

		} else {
			expireToken := time.Now().Add(time.Hour * 666).Unix()
			signedToken, err := models.GenerateToken(user, ctx.Host(), expireToken)
			if err != nil {
				res := response.InternalServerError("", err.Error())
				ctx.JSON(res.StatusCode(), res.Struct())
				return
			}

			token, err = models.SaveToken(signedToken, expireToken, user, scope.(app.RequestScope).GetTx())
			if err != nil {
				res := response.InternalServerError("", err.Error())
				ctx.JSON(res.StatusCode(), res.Struct())
				return
			}
		}

		ctx.JSON(iris.StatusOK, iris.Map{"token": token.Token, "refresh_token": token.RefreshToken, "expires at": token.ExpiresAt})
		return
	}

	res := response.Unauthorized("")
	ctx.JSON(res.StatusCode(), res.Struct())
}
