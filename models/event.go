package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"errors"
)

type Event struct {
	gorm.Model
	Title string `json:"title"`
	Body string `json:"body"`
	Image Image `json:"image_id"`
	Team Team `json:"team_id"`
	ImagePath string `json:"image_path"`
	AccountId int `json:"account_id"`
	EventDate int64 `json:"event_date"`
}

// Migrate Event structure
func EventMigrate() error {
	db, err := gorm.Open("mysql", "root:@/planner?charset=utf8&parseTime=True&loc=Local")
	if !db.HasTable(&Event{}) {
		if err != nil {
			return errors.New("Problem with connect to database")
		}
		db.AutoMigrate(&Event{})
		defer db.Close()
	}

	return nil
}
