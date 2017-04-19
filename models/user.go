package models

import (
	"github.com/jinzhu/gorm"
	"errors"
)

type User struct {
	gorm.Model
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}

// Migrate User structure
func UserMigrate() error {
	db, err := gorm.Open("mysql", "root:@/planner?charset=utf8&parseTime=True&loc=Local")
	if !db.HasTable(&User{}) {
		if err != nil {
			return errors.New("Problem with connect to database")
		}
		db.AutoMigrate(&User{})
		defer db.Close()
	}

	return nil
}
