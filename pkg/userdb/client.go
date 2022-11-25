// https://github.com/iammukeshm/jwt-authentication-golang

package userdb //TODO: Refactor this and mysql package to not collide

import (
	"log"
	"strings"

	"github.com/mikejk8s/gmud/pkg/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Instance *gorm.DB
var dbError error

func Connect(userDBusername string, userDBpassword string, userDBhostname string, userDBdbname string) {
	Instance, dbError = gorm.Open(mysql.Open(userDBusername+":"+userDBpassword+"@"+"tcp"+userDBhostname+"/"+userDBdbname), &gorm.Config{})
	if dbError != nil && strings.Contains(dbError.Error(), "Unknown database 'users'") {
		Instance, dbError = gorm.Open(mysql.Open(userDBusername+":"+userDBpassword+"@"+"tcp"+userDBhostname+"/"), &gorm.Config{})
		Instance.Exec("CREATE DATABASE IF NOT EXISTS " + userDBdbname)
		Instance.Exec(`CREATE TABLE IF NOT EXISTS users(
    	id            integer unsigned null,
    	created_at    datetime         null,
   	 	updated_at    datetime         null,
    	deleted_at    datetime         null,
    	name          varchar(255)     null,
    	password_hash varchar(255)     null,
    	remember_hash varchar(255)     null
		);`)
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
