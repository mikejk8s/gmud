package mysqlpkg

import (
	"context"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func (s *SqlConn) CreateCharacterTable() error {
	query := `CREATE TABLE IF NOT EXISTS characters (
		id BIGINT UNIQUE NOT NULL PRIMARY KEY,
		name VARCHAR(30) UNIQUE NOT NULL,
		class VARCHAR(15) NOT NULL,
		race VARCHAR(15) NOT NULL DEFAULT 'HUMAN',
		level INT(3) NOT NULL DEFAULT '1',
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		alive BOOLEAN NOT NULL DEFAULT '1',
        characterowner VARCHAR(20) NOT NULL DEFAULT 'player'
	) ENGINE=INNODB;`

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	res, err := s.DB.ExecContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when creating Character table", err)
		panic(err)
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when getting rows affected", err)
		panic(err)
		return err
	}
	log.Printf("Rows affected when creating table: %d", rows)
	return nil
}
