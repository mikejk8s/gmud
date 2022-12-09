package mysqlpkg

import (
	"github.com/mikejk8s/gmud/pkg/userdb"
	"log"
)

func (s *SqlConn) CreateUsersTable() {
	s.DB.Exec(`CREATE TABLE IF NOT EXISTS userstable(
    	id            integer unsigned null,
    	created_at    datetime         null,
   	 	updated_at    datetime         null,
    	deleted_at    datetime         null,
    	name          varchar(255)     null,
    	password_hash varchar(255)     null,
    	remember_hash varchar(255)     null
		);`)
}

// CreateNewUser is self-explanatory by its name.
//
// Created_at and updated_at are set to the current time of the database.
//
// LoginReq is a struct that contains the user's name, password and email, that is sent from signup server mostly, or you can set one up yourself.
func (s *SqlConn) CreateNewUser(userInfo LoginReq) error {
	stmt, err := s.DB.Prepare("INSERT INTO users.users (created_at,updated_at,deleted_at,name, username,email,password) VALUES (CURRENT_TIMESTAMP,CURRENT_TIMESTAMP,null,?,?,?,?)")
	if err != nil {
		log.Println("Error", err.Error())
	}
	_, err = stmt.Query(userInfo.Name, userInfo.Name, userInfo.Email, userInfo.Password)
	if err != nil {
		return err
	}
	return nil
}
func Migration() {
	userdb.Connect(Username, Password, Hostname, "users")
	userdb.Migrate()
}
