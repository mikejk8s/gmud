package mysqlpkg

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
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
	db, err := sql.Open("mysql", Username+":"+Password+"@tcp"+Hostname+"/"+dbname+"?parseTime=true")
	if err != nil {
		log.Println("Error", err.Error())
		return err
	}
	conn.DB = db
	return nil
}

var RunningOnDocker = false

// these are self-explanatory, change them to your own database credentials when you are running them locally.
//
// will change it to work with docker-composer environment variables later.
var (
	Username = "cansu"
	Password = "1234"
	//hostname = "docker.for.mac.localhost:3306"
	Hostname = "(127.0.0.1:3306)"
)
