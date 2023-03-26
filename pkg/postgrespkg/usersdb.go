package postgrespkg

import (
	"log"

	"github.com/mikejk8s/gmud/pkg/userdb"
)

func (s *SqlConn) CreateUsersTable() {
	s.DB.Exec(`CREATE TABLE IF NOT EXISTS users (
		id              SERIAL PRIMARY KEY,
		created_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at      TIMESTAMP,
		deleted_at      TIMESTAMP,
		name            VARCHAR(255),
		password_hash   VARCHAR(255),
		remember_hash   VARCHAR(255)
	);`)
}

// CreateNewUser is self-explanatory by its name.
//
// Created_at and updated_at are set to the current time of the database.
//
// LoginReq is a struct that contains the user's name, password and email, that is sent from signup server mostly, or you can set one up yourself.
func (s *SqlConn) CreateNewUser(userInfo LoginReq) error {
	stmt, err := s.DB.Prepare("INSERT INTO users (created_at, updated_at, deleted_at, name, username, email, password) VALUES (CURRENT_TIMESTAMP, NULL, NULL, ?, ?, ?, ?)")
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
