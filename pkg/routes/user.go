package routes

import (
	"context"
	"database/sql"
	"github.com/mikejk8s/gmud/pkg/userdb"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

func ConnectUserDB() (*sql.DB, error) {
	// Initialize Database
	characterDB, err := gorm.Open(mysql.Open("root:1234@tcp(127.0.0.1:3306)/"), &gorm.Config{})
	_, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	characterDB.Exec("CREATE SCHEMA IF NOT EXISTS " + "users")
	if err != nil {
		log.Printf("Error %s when creating DB\n", err)
		panic(err)
	}

	return characterDB.DB()
}

func Migration() {
	userdb.Connect("root:1234@tcp(127.0.0.1:3306)/users?parseTime=true")
	userdb.Migrate()
}
