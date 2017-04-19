package models

import (
	"github.com/jinzhu/gorm"
	"errors"
)

type Image struct {
	gorm.Model
	Path string `json:"path"`
}

// Migrate Image structure
func ImageMigrate() error {
	db, err := gorm.Open("mysql", "root:@/planner?charset=utf8&parseTime=True&loc=Local")
	if !db.HasTable(&Image{}) {
		if err != nil {
			return errors.New("Problem with connect to database")
		}
		db.AutoMigrate(&Image{})
		defer db.Close()
	}

	return nil
}