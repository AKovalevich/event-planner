package models

import (
	"github.com/AKovalevich/event-planner/utils"

	"errors"
	"github.com/jinzhu/gorm"
)

type Image struct {
	gorm.Model
	Path string `json:"path"`
}

// Migrate Image structure
func ImageMigrate() error {
	db, err := utils.GetDB()
	if !db.HasTable(&Image{}) {
		if err != nil {
			return errors.New("Problem with connect to database")
		}
		db.AutoMigrate(&Image{})
		defer db.Close()
	}

	return nil
}