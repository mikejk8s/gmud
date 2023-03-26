package userdb //TODO: Refactor this and mysql package to not collide

import (
	"log"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/mikejk8s/gmud/pkg/models"
)

var Instance *gorm.DB
var dbError error

func Connect(userDBusername string, userDBpassword string, userDBhostname string, userDBdbname string) {
	Instance, dbError = gorm.Open(postgres.Open("host="+userDBhostname+" user="+userDBusername+" password="+userDBpassword+" dbname="+userDBdbname+" port=5432 sslmode=disable"), &gorm.Config{})
	if dbError != nil && strings.Contains(dbError.Error(), "Unknown database 'users'") {
		Instance, dbError = gorm.Open(postgres.Open("host="+userDBhostname+" user="+userDBusername+" password="+userDBpassword+" port=5432 sslmode=disable"), &gorm.Config{})
		Instance.Exec("CREATE DATABASE " + userDBdbname)
		Instance.Exec(`CREATE TABLE IF NOT EXISTS users(
			id            serial primary key,
			created_at    timestamp         null,
			updated_at    timestamp         null,
			deleted_at    timestamp         null,
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
