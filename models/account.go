package models

import (
	"github.com/AKovalevich/event-planner/utils"

	_"github.com/jinzhu/gorm"
	"errors"
	"github.com/asaskevich/govalidator"
	"time"
	"github.com/jinzhu/gorm"
)

type Account struct {
	BaseModel `valid:"-"`
	User string `json:"user" valid:"-"`
	Email string `json:"email" valid:"-"`
	Teams []Team `json:"teams" gorm:"many2many:team_account" valid:"-"`
}

// Migrate Team structure
func AccountMigrate() error {
	db, err := utils.GetDB()
	if !db.HasTable(&Account{}) {
		if err != nil {
			return errors.New("Problem with connect to database")
		}
		db.AutoMigrate(&Account{})
		defer db.Close()
	}

	return nil
}

//
func CreateAccount(account *Account, tx interface{}) (*Account, error) {
	// prepare values for default fields
	account.UpdatedAt = time.Now().UTC().UnixNano() / int64(time.Second)
	account.CreatedAt = time.Now().UTC().UnixNano() / int64(time.Second)

	// base structure validation
	if _, err := govalidator.ValidateStruct(account); err != nil {
		println(err.Error())
		return account, err
	}

	// create new event
	tx.(*gorm.DB).Create(&account)

	// in case if something is wrong with MySQL insert operation
	if account.ID == 0 {
		return account, errors.New("Service unavalible")
	}

	return account, nil
}
