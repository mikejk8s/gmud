package postgrespkg

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	m "github.com/mikejk8s/gmud/pkg/models"
)

func (s *SqlConn) Close() error {
	if err := s.DB.Close(); err != nil {
		return fmt.Errorf("error closing SQL connection: %w", err)
	}
	return nil
}

func (s *SqlConn) GetCharactersByOwner(owner string) ([]*m.Character, error) {
	rows, err := s.DB.Query("SELECT * FROM characters WHERE characterowner = $1", owner)
	if err != nil {
		return nil, fmt.Errorf("error querying characters: %w", err)
	}
	defer rows.Close()

	var characters []*m.Character
	for rows.Next() {
		var character m.Character
		if err := rows.Scan(&character.ID, &character.Name, &character.Class, &character.Race, &character.Level, &character.CreatedAt, &character.Alive, &character.CharacterOwner); err != nil {
			return nil, fmt.Errorf("error scanning row: %w", err)
		}
		character.CharacterOwner = owner
		characters = append(characters, &character)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error getting row errors: %w", err)
	}
	return characters, nil
}

func (s *SqlConn) AddCharacter(Character m.Character) {
	insert, err := s.DB.Query(
		"INSERT INTO characters (name,id,class,level,race,characterowner) VALUES ($1,$2,$3,$4,$5,$6)",
		Character.Name, Character.ID, Character.Class, Character.Level, Character.Race, Character.CharacterOwner)
	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	defer func(insert *sql.Rows) {
		err := insert.Close()
		if err != nil {
			log.Println(err)
		}
	}(insert)
}

func (s *SqlConn) DeleteCharacter(Character m.Character) {
	delete, err := s.DB.Query(
		"DELETE FROM characters WHERE id = $1", Character.ID)
	// if there is an error deleting, handle it
	if err != nil {
		panic(err.Error())
	}
	defer func(delete *sql.Rows) {
		err := delete.Close()
		if err != nil {
			log.Println(err)
		}
	}(delete)
}

func (s *SqlConn) CreateCharacterTable() error {
	query := `CREATE TABLE IF NOT EXISTS characters (
		id BIGINT UNIQUE NOT NULL PRIMARY KEY,
		name VARCHAR(30) UNIQUE NOT NULL,
		class VARCHAR(15) NOT NULL,
		race VARCHAR(15) NOT NULL DEFAULT 'HUMAN',
		level INT NOT NULL DEFAULT '1',
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		alive BOOLEAN NOT NULL DEFAULT '1',
		characterowner VARCHAR(20) NOT NULL DEFAULT 'player'
	)`
	if _, err := s.DB.Exec(query); err != nil {
		return fmt.Errorf("error creating Character table: %w", err)
	}

	log.Printf("Character table created successfully.")
	return nil
}