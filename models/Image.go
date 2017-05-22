package models

import (
	"github.com/AKovalevich/event-planner/utils"

	_ "github.com/jinzhu/gorm"
	"errors"

)

type Image struct {
	BaseModel
	Path string `json:"path"`
	Team Team `json:"team"`
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