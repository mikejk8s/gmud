// https://github.com/iammukeshm/jwt-authentication-golang

package userdb //TODO: Refactor this and mysql package to not collide

import (
	"log"

	"github.com/mikejk8s/gmud/pkg/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Instance *gorm.DB
var dbError error

func Connect(connectionString string) {
	Instance, dbError = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if dbError != nil {
		panic("Cannot connect to DB")
	}
	log.Println("Connected to Database!")
}

func Migrate() {
	err := Instance.AutoMigrate(&models.User{})
	if err != nil {
		panic(err)
	}
	log.Println("Database Migration Completed!")
}
