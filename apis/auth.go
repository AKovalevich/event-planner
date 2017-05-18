package apis

import (
	"github.com/AKovalevich/event-planner/response"
	"github.com/AKovalevich/event-planner/models"
	"github.com/AKovalevich/event-planner/app"

	"github.com/dgrijalva/jwt-go"
	"gopkg.in/kataras/iris.v6"
	"encoding/hex"
	"crypto/md5"
	"time"
)


func Auth(ctx *iris.Context) {
	credentials := models.Credentials{}
	err := ctx.ReadJSON(&credentials)
	if err != nil {
		res := response.InternalServerError("", err.Error())
		ctx.JSON(res.StatusCode(), res.Struct())
		return
	}

	// prepare password hash
	hash := md5.New()
	hash.Write([]byte(credentials.Password))
	hashedPassword := hex.EncodeToString(hash.Sum(nil))

	// try to load user by password and email.
	user, err := models.GetUserByEmail(credentials.Email)
	if err != nil {
		res := response.NotFound("User")
		ctx.JSON(res.StatusCode(), res.Struct())
		return
	}
	if credentials.Email == "test@test.com" && credentials.Password == hashedPassword { // "test" - "098f6bcd4621d373cade4e832627b4f6"
		// expires the token and cookie in 1 hour
		expireToken := time.Now().Add(time.Hour * 666).Unix()

		claims := models.Claims {
			jwt.StandardClaims {
				ExpiresAt: expireToken,
				Issuer:    "localhost:9000",
			},
			*user,
		}

		// create the token using your claims
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		secret := app.Config().Secret

		// signs the token with a secret.
		signedToken, err := token.SignedString([]byte(secret))

		if err != nil {
			res := response.InternalServerError("", err.Error())
			ctx.JSON(res.StatusCode(), res.Struct())
			return
		}

		ctx.JSON(iris.StatusOK, signedToken)
	}

	res := response.Unauthorized("")
	ctx.JSON(res.StatusCode(), res.Struct())
	return
}
