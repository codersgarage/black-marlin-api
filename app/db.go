package app

import (
	"fmt"
	"github.com/codersgarage/black-marlin-api/config"
	"github.com/codersgarage/black-marlin-api/log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var instance *gorm.DB

func ConnectDB() error {
	db, err := gorm.Open("postgres",
		fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=disable",
			config.DB().Username, config.DB().Password,
			config.DB().Host, config.DB().Port, config.DB().Name))
	if err != nil {
		return err
	}

	db.LogMode(true)
	db.SetLogger(log.Log())

	instance = db
	return nil
}

func DB() *gorm.DB {
	return instance
}
