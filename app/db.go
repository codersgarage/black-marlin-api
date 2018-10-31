package app

import (
	"fmt"

	"github.com/codersgarage/black-marlin-api/config"
	"github.com/codersgarage/black-marlin-api/log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var instance *gorm.DB

func ConnectDB() error {
	db, err := gorm.Open("mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
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
