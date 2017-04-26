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
