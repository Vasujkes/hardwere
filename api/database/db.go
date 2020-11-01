package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"./config"
)

func Connect() *gorm.DB {
	db, err := gorm.Open("mysql", config.BuildDSN())
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return db
}
