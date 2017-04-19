package models

import (
	"github.com/jinzhu/gorm"
	"errors"
)

type Account struct {
	gorm.Model
	User string `json:"user"`
	Email string `json:"email"`
}

// Migrate Team structure
func AccountMigrate() error {
	db, err := gorm.Open("mysql", "root:@/planner?charset=utf8&parseTime=True&loc=Local")
	if !db.HasTable(&Account{}) {
		if err != nil {
			return errors.New("Problem with connect to database")
		}
		db.AutoMigrate(&Account{})
		defer db.Close()
	}

	return nil
}
