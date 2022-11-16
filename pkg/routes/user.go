package routes

import (
	"context"
	"database/sql"
	"github.com/mikejk8s/gmud/logger"
	"github.com/mikejk8s/gmud/pkg/userdb"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func ConnectUserDB() (*sql.DB, error) {
	// Initialize Database
	characterDB, err := gorm.Open(mysql.Open("root:1234@tcp(127.0.0.1:3306)/"), &gorm.Config{})
	_, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	characterDB.Exec("CREATE SCHEMA IF NOT EXISTS " + "users")
	characterDBLogger := logger.GetNewLogger()
	characterDBLogger.AssignOutput("characterDB", "./app/characterdbconn")
	if err != nil {
		characterDBLogger.LogUtil.Errorln("Error %s when opening DB", err)
		panic(err)
	} else {
		characterDBLogger.LogUtil.Infoln("Connected to Characters DB")
	}

	return characterDB.DB()
}

func Migration() {
	userdb.Connect("root:1234@tcp(127.0.0.1:3306)/users?parseTime=true")
	userdb.Migrate()
}
