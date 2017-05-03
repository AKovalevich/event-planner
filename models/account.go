package models

import (
	"github.com/AKovalevich/event-planner/utils"

	"errors"
	_"github.com/jinzhu/gorm"
)

type Account struct {
	BaseModel
	User string `json:"user"`
	Email string `json:"email"`
}

// Migrate Team structure
func AccountMigrate() error {
	db, err := utils.GetDB()
	if !db.HasTable(&Account{}) {
		if err != nil {
			return errors.New("Problem with connect to database")
		}
		db.AutoMigrate(&Account{})
		defer db.Close()
	}

	return nil
}
