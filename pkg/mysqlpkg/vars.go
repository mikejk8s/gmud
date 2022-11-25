package mysqlpkg

import (
	"database/sql"
	"log"
)

// Dont forget to call SqlConn.CloseConn when you are done with the connection.
type SqlConn struct {
	DB *sql.DB
}

// GetSQLConn attaches a new sql connection to the SqlConn struct.
func (conn *SqlConn) GetSQLConn(dbname string) error {
	db, err := sql.Open("mysql", username+":"+password+"@tcp"+hostname+"/"+dbname+"?parseTime=true")
	if err != nil {
		log.Println("Error", err.Error())
		return err
	}
	conn.DB = db
	return nil
}

const (
	username = "cansu"
	password = "1234"
	//hostname = "docker.for.mac.localhost:3306"
	hostname = "(127.0.0.1:3306)"
)
