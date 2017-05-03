package models

import (
	"github.com/AKovalevich/event-planner/utils"

	"errors"
	_ "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm"
	"github.com/asaskevich/govalidator"
	"time"
)

type Team struct {
	BaseModel `valid:"optional"`
	Name string `json:"name" valid:"alphanum,required" gorm:"type:varchar(100)"`
	Description string `json:"description" valid:"-" gorm:"size:255"`
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
	// Base structure validation
	if _, err := govalidator.ValidateStruct(team); err != nil {
		return team, err
	}

	// Check that we still don't have a team with the same name
	// @TODO need to move to custom validation tag
	db, err := utils.GetDB()
	if err != nil {
		return team, err
	}
	existingTeam := &Team{}
	db.Where("name = ?", team.Name).First(&existingTeam)
	if len(existingTeam.Name) > 0 {
		return team, errors.New("Team already exist")
	}

	// Prepare values for default fields
	team.UpdatedAt = time.Now().UTC().UnixNano() / int64(time.Second)
	team.CreatedAt = time.Now().UTC().UnixNano() / int64(time.Second)

	// Create new team
	db.Create(&team)

	// In case if something is wrong with MySQL insert operation
	if team.ID == 0 {
		return team, errors.New("Service unavalible")
	}

	return team, nil
}
