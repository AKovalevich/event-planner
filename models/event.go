package models

import (
	"github.com/AKovalevich/event-planner/utils"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm"
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/asaskevich/govalidator"
	"time"
)

type Event struct {
	BaseModel `valid:"-"`
	Title string `json:"title" valid:"-"`
	Body string `json:"body" valid:"-"`
	Image Image `json:"image_id" valid:"-"`
	TeamID uint `json:"team_id" valid:"required"`
	UserID uint `json:"user_id" valid:"required"`
	ImagePath string `json:"image_path" valid:"-"`
	AccountId uint `json:"account_id" valid:"required"`
	EventDate int64 `json:"event_date" valid:"-"`
	Status bool `json:"status" valid:"required"`
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

//
func GetTeamEvents(teamId uint, tx interface{}) []Event {
	var events = []Event{}
	tx.(*gorm.DB).Where("team_id = ?", teamId).Find(&events)

	return events
}

//
func CreateEvent(event *Event, tx interface{}) (*Event, error) {
	// base structure validation
	if _, err := govalidator.ValidateStruct(event); err != nil {
		return event, err
	}

	// prepare values for default fields
	event.UpdatedAt = time.Now().UTC().UnixNano() / int64(time.Second)
	event.CreatedAt = time.Now().UTC().UnixNano() / int64(time.Second)

	// create new event
	tx.(*gorm.DB).Create(&event)

	// in case if something is wrong with MySQL insert operation
	if event.ID == 0 {
		return event, errors.New("Service unavalible")
	}

	return event, nil
}
