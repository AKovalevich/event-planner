package utils

import (
	"github.com/AKovalevich/event-planner/app"

	"github.com/jinzhu/gorm"
	"time"
)

//
func GetDB() (*gorm.DB, error) {
	// Try ping to check connection
	attempts := 10
	for i := 0; i < attempts; i++ {
		db, err := gorm.Open("mysql", app.Config().DataBase.User+ ":@/" + app.Config().DataBase.Name + "?charset=utf8&parseTime=True&loc=Local")
		if err == nil {
			db.Close()
			break
		}
		println("Connect to database failed. Try to reconnect to database...")
		println("DB user: " + app.Config().DataBase.User)
		println("DB name: " + app.Config().DataBase.Name)
		time.Sleep(5 * time.Second)
	}

	db, err := gorm.Open("mysql", app.Config().DataBase.User + ":@/" + app.Config().DataBase.Name + "?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return nil, err
	}

	db.LogMode(true)

	return db, nil
}
