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
	db, err := gorm.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		config.DB().Host, config.DB().Port, config.DB().Username, config.DB().Name, config.DB().Password))
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
