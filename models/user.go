package models

import (
	"github.com/AKovalevich/event-planner/utils"

	"github.com/asaskevich/govalidator"
	"errors"
	"time"
	"github.com/jinzhu/gorm"
)

//
type User struct {
	BaseModel `valid:"optional"`
	Name string `json:"name" valid:"-"`
	Email string `json:"email" valid:"required"`
	Password string `json:"password" valid:"required"`
	Teams []Team `json:"teams" gorm:"many2many:team_user" valid:"-"`
	TokenID uint `json:"token_id" valid:"-"`
	Events []Event
}

// Migrate User structure
func UserMigrate() error {
	db, err := utils.GetDB()
	if !db.HasTable(&User{}) {
		if err != nil {
			return errors.New("Problem with connect to database")
		}
		db.AutoMigrate(&User{})
		defer db.Close()
	}

	return nil
}

//
func GetUserByEmail(email string, tx interface{}) (*User, error) {
	user := &User{}

	tx.(*gorm.DB).Where("email = ?", email).First(&user);

	return user, nil
}

//
func CreateUser(user *User, tx interface{}) (*User, error) {
	// Base structure validation
	if _, err := govalidator.ValidateStruct(user); err != nil {
		return user, err
	}

	// Check that we still don't have a team with the same name
	// @TODO need to move to custom validation tag
	existingUser, err := GetUserByEmail(user.Email, tx)
	if err != nil {
		return user, err
	}
	if len(existingUser.Email) > 0 {
		return user, errors.New("User already exist")
	}

	// Prepare values for default fields
	user.UpdatedAt = time.Now().UTC().UnixNano() / int64(time.Second)
	user.CreatedAt = time.Now().UTC().UnixNano() / int64(time.Second)

	// Create new team
	tx.(*gorm.DB).Create(&user)

	// In case if something is wrong with MySQL insert operation
	if user.ID == 0 {
		return user, errors.New("Service unavalible")
	}

	return user, nil
}

//
func (user *User) LoadUserAssociations(tx interface{}) error {
	tx.(*gorm.DB).Preload("Teams").Preload("Role").Find(&user)

	return nil
}
