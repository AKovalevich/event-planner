package models

import (
	"github.com/AKovalevich/event-planner/app"
	"github.com/AKovalevich/event-planner/utils"
	"time"
	"errors"
	"github.com/dgrijalva/jwt-go"
)

type Token struct {
	BaseModel
	Token string `json:"token" gorm:"size:512"`
	RefreshToken string `json:"refresh_token" gorm:"size:255"`
	User User `json:"user"`
	ExpiresAt int64 `json:"expires_at"`
}

// migrate Team structure
func TokenMigrate() error {
	db, err := utils.GetDB()
	if !db.HasTable(&Token{}) {
		if err != nil {
			return errors.New("Problem with connect to database")
		}
		db.AutoMigrate(&Token{})
		defer db.Close()
	}

	return nil
}

//
func (token *Token) Expired() bool {
	currentTime := int64(time.Now().Unix())
	if currentTime >= token.ExpiresAt {
		return true
	} else {
		return false
	}
}

//
func (token *Token) UpdateToken(user *User, issuer string) error {
	db, err := utils.GetDB()
	defer db.Close()
	if err != nil {
		return err
	}

	expireToken := time.Now().Add(time.Hour * 666).Unix()
	newToken, err := GenerateToken(user, issuer, expireToken)
	if err != nil {
		return err
	}
	newRefreshToken := GenerateRefreshToken(newToken, app.Config().Secret)
	token.RefreshToken = newRefreshToken
	token.Token = newToken
	token.UpdatedAt = time.Now().UTC().UnixNano() / int64(time.Second)
	db.Model(&token).Omit("id", "created_at", "deleted").Update(token)

	return nil
}

//
func SaveToken(tokenString string, expires_at int64, user *User) (*Token, error) {
	db, err := utils.GetDB()
	defer db.Close()
	if err != nil {
		return &Token{}, err
	}

	secret := app.Config().Secret
	var token = &Token{
		Token: tokenString,
		RefreshToken: GenerateRefreshToken(tokenString, secret),
		ExpiresAt: expires_at,
		User: *user,
	}

	token.CreatedAt = time.Now().UTC().UnixNano() / int64(time.Second)
	token.UpdatedAt = time.Now().UTC().UnixNano() / int64(time.Second)
	token.Deleted = false

	if err := db.Save(token).Error; err != nil {
		return &Token{}, err
	}

	return token, nil
}

//
func LoadTokenQuery(query interface{}) (*Token, error) {
	var token = &Token{}
	db, err := utils.GetDB()
	defer db.Close()
	if err != nil {
		return token, err
	}

	db.Where(query).First(&token)

	return token, nil
}

//
func LoadToken(tokenString string) (*Token, error) {
	var token = &Token{}

	db, err := utils.GetDB()
	defer db.Close()
	if err != nil {
		return token, err
	}

	db.Where("token = ?", tokenString).First(&token)

	return token, nil
}

//
func DeleteToken(tokenString string) error {
	db, err := utils.GetDB()
	defer db.Close()
	if err != nil {
		return err
	}

	if err = db.Where("token = ?", tokenString).Delete(Token{}).Error; err != nil {
		return err
	}

	return nil
}

//func LoadUserToken(user User) (Token, error) {
//	db, err := utils.GetDB()
//	defer db.Close()
//	if err != nil {
//		return err
//	}
//
//	db.Where(" = ?", tokenString).First(&token)
//}

//
func GenerateRefreshToken(token string, secret string) (string) {
	return utils.HashMd5(token + secret)
}

//
func GenerateToken(user *User, issuer string, expire int64) (string, error) {
	claims := Claims {
		jwt.StandardClaims {
			ExpiresAt: expire,
			Issuer:    issuer,
		},
		*user,
	}
	// create the token using your claims
	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := app.Config().Secret

	// signs the token with a secret.
	signedToken, err := rawToken.SignedString([]byte(secret))

	return signedToken, err
}