package postgrespkg

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

// LoginReq is used for handling sign-up requests from the webpage server.
//
// LoginReq can be used for creating a new user with SqlConn model.
type LoginReq struct {
	Name     string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// Dont forget to call SqlConn.CloseConn when you are done with the connection.
type SqlConn struct {
	DB *sql.DB
}

// GetSQLConn attaches a new sql connection to the SqlConn struct.
func (conn *SqlConn) GetSQLConn(dbname string) error {
	dsn := fmt.Sprintf("%s:%s@%s/postgres?parseTime=true&sslmode=%s", Username, Password, Hostname, SSLMode)
	println(dsn)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Println("Error:", err.Error())
		return err
	}
	conn.DB = db

	query := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s;", dbname)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if _, err := conn.DB.ExecContext(ctx, query); err != nil {
		return fmt.Errorf("error creating Character database: %w", err)
	}

	return nil
}

var RunningOnDocker = false

// these are self-explanatory, change them to your own database credentials when you are running them locally.
//
// will change it to work with docker-composer environment variables later.
var (
	Username = "gmud"
	Password = "gmud"
	//hostname = "docker.for.mac.localhost:3306"
	Hostname = "127.0.0.1:5432"
	SSLMode  = "disable"
)
