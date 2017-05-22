package models

import (
	"github.com/AKovalevich/event-planner/utils"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm"
	"errors"
)

type Event struct {
	BaseModel
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
	db, err := utils.GetDB()
	if !db.HasTable(&Event{}) {
		if err != nil {
			return errors.New("Problem with connect to database")
		}
		db.AutoMigrate(&Event{})
		defer db.Close()
	}

	return nil
}
