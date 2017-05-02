package models

import (
	"github.com/AKovalevich/event-planner/utils"

	"errors"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

)

type Team struct {
	gorm.Model
	Name string `json:"name" validate:"required" gorm:"type:varchar(100)"`
	Description string `json:"description" gorm:"size:255"`
}

// Migrate Team structure
func TeamMigrate() error {
	db, err := utils.GetDB()
	if !db.HasTable(&Team{}) {
		if err != nil {
			return errors.New("Problem with connect to database")
		}
		db.AutoMigrate(&Team{})
		defer db.Close()
	}

	return nil
}

func CreateTeam(team *Team) (*Team, error) {
	validator := utils.GetValidator()
	err := validator.Struct(team)
	return team, err
}
