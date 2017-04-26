package utils

import (
	"github.com/AKovalevich/event-planner/config"

	"github.com/jinzhu/gorm"
	"time"
)

func GetDB() (*gorm.DB, error) {
	appConfig := config.Get()
	// Try ping to check connection
	attempts := 10
	for i := 0; i < attempts; i++ {
		db, err := gorm.Open("mysql", appConfig.DataBase.User+ ":@/" + appConfig.DataBase.Name + "?charset=utf8&parseTime=True&loc=Local")
		if err == nil {
			db.Close()
			break
		}
		println("Connect to database failed. Try to reconnect to database...")
		time.Sleep(5 * time.Second)
	}

	db, err := gorm.Open("mysql", appConfig.DataBase.User+ ":@/" + appConfig.DataBase.Name + "?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return nil, err
	}

	db.LogMode(true)

	return db, nil
}
