package models

import (
	"github.com/AKovalevich/event-planner/utils"

	_"github.com/jinzhu/gorm"
	"errors"
)

type Account struct {
	BaseModel
	User string `json:"user"`
	Email string `json:"email"`
	Teams []Team `json:"teams" gorm:"many2many:team_account"`
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
