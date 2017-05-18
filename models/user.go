package models

import (
	"github.com/AKovalevich/event-planner/utils"

	"errors"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
	Teams []Team `json:"teams" gorm:"many2many:team_user"`
}

// Migrate User structure
func UserMigrate() error {
	db, err := utils.GetDB()
	if !db.HasTable(&User{}) {
		if err != nil {
			return errors.New("Problem with connect to database")
		}
		db.AutoMigrate(&User{})
		defer db.Close()
	}

	return nil
}

func GetUserByEmail(email string) (*User, error) {
	user := &User{}

	db, err := utils.GetDB()
	if err != nil {
		return user, err
	}

	if err:= db.Where("email = ?", email).First(&user).Error; err != nil {
		return user, err
	}

	if user.Email == "" {
		return user, errors.New("Can't load user")
	}

	return user, nil
}
