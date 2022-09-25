// https://github.com/iammukeshm/jwt-authentication-golang

package userdb //TODO: Refactor this and mysql package to not collide

import (
	"github.com/mikejk8s/gmud/pkg/models"
	"log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Instance *gorm.DB
var dbError error

func Connect(connectionString string) () {
	Instance, dbError = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if dbError != nil {
		log.Fatal(dbError)
		panic("Cannot connect to DB")
	}
	log.Println("Connected to Database!")
}

func Migrate() {
	Instance.AutoMigrate(&models.User{})
	log.Println("Database Migration Completed!")
}

func ConnectUserDB() {
	// Initialize Database
	Connect("user:password@tcp(localhost:3307)/users?parseTime=true")
	Migrate()
}