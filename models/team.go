package models

import (
	"github.com/AKovalevich/event-planner/utils"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/asaskevich/govalidator"
	_ "github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm"
	"errors"
	"time"
)

//
type Team struct {
	BaseModel `valid:"optional"`
	Name string `json:"name" valid:"required" gorm:"type:varchar(100)"`
	Description string `json:"description" valid:"-" gorm:"size:255"`
	Users []User `json:"users, omitempty" valid:"-" gorm:"many2many:team_user"`
	Accounts []Account `json:"accounts, omitempty" valid:"-" gorm:"many2many:team_account"`
	Status bool `json:"status" valid:"-"`
	Events []Event
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

func (team *Team) HasUser(userID uint, tx interface{}) (bool) {
	var existingTeam = &Team{}
	if err := tx.(*gorm.DB).Preload("Users", "id = ?", userID).Where("id = ?", team.ID).First(&existingTeam).Error; err != nil {
		return false
	}

	if len(existingTeam.Users) >= 1 {
		return true
	}

	return false
}

func (team *Team) HasAccount(accountID uint, tx interface{}) (bool) {
	var existingTeam = &Team{}
	if err := tx.(*gorm.DB).Preload("Accounts", "id = ?", accountID).Where("id = ?", team.ID).First(&existingTeam).Error; err != nil {
		return false
	}

	if len(existingTeam.Accounts) >= 1 {
		return true
	}

	return false
}

//
func CreateTeam(team *Team, tx interface{}) (*Team, error) {
	// Base structure validation
	if _, err := govalidator.ValidateStruct(team); err != nil {
		return team, err
	}

	// Check that we still don't have a team with the same name
	// @TODO need to move to custom validation tag
	existingTeam := &Team{}
	tx.(*gorm.DB).Where("name = ?", team.Name).First(&existingTeam)
	if len(existingTeam.Name) > 0 {
		return team, errors.New("Team already exist")
	}

	// Prepare values for default fields
	team.UpdatedAt = time.Now().UTC().UnixNano() / int64(time.Second)
	team.CreatedAt = time.Now().UTC().UnixNano() / int64(time.Second)

	// Create new team
	tx.(*gorm.DB).Create(&team)

	// In case if something is wrong with MySQL insert operation
	if team.ID == 0 {
		return team, errors.New("Service unavalible")
	}

	return team, nil
}

//
func GetTeam(id uint, tx interface{}) (*Team, error) {
	team := &Team{}

	if err := tx.(*gorm.DB).Where("id = ?", id).First(&team).Error; err != nil {
		return team, err
	}

	if team.ID == 0 {
		return team, errors.New("Team " + string(id) + " not found")
	}

	return team, nil
}

func GetTeamUsers(id string, tx interface{}) ([]User, error) {
	var users = []User{}
	var team = &Team{}

	if err := tx.(*gorm.DB).Preload("Users").Where("id = ?", id).First(&team).Error; err != nil {
		return users, err
	}

	return team.Users, nil
}
