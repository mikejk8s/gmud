package mysqlpkg

import (
	"context"
	"fmt"
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
	)`
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if _, err := s.DB.ExecContext(ctx, query); err != nil {
		return fmt.Errorf("error creating Character table: %w", err)
	}

	log.Printf("Character table created successfully.")
	return nil
}
