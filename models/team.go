package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"errors"
)

type Team struct {
	gorm.Model
	Name string `json:"name" gorm:"type:varchar(100)"`
	Description string `json:"description" gorm:"size:255"`
}

// Migrate Team structure.
func TeamMigrate() error {
	db, err := gorm.Open("mysql", "root:@/planner?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return errors.New("Problem with connect to database")
	}
	db.AutoMigrate(&Team{})
	defer db.Close()

	return nil
}
