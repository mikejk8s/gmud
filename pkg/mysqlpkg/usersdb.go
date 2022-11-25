package mysqlpkg

import (
	"github.com/mikejk8s/gmud/pkg/userdb"
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

func Migration() {
	userdb.Connect(username, password, hostname, "users")
	userdb.Migrate()
}
