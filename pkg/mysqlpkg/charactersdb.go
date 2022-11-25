package mysqlpkg

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	//"github.com/mikejk8s/gmud"
	"fmt"

	//cr "github.com/mikejk8s/gmud/pkg/charactersroutes"
	m "github.com/mikejk8s/gmud/pkg/models"
)

// func GetCharacters(code string) []m.Character {
// 	DB, err := sql.Open("mysql", username+":"+password+"@tcp(127.0.0.1:3306)/"+dbname)
// 	char:= &m.Character{}
// 	if err != nil {
// 		fmt.Println("Error", err.Error())
// 		return nil
// 		{
// 	defer DB.Close()
// 	results, err := DB.Query("SELECT * FROM characters")
// 	if err != nil {
// 		fmt.Println("Err", err.Error())
// 		return nil
// 	}
// 	defer results.Close()
// 	for result.Next()

func (s *SqlConn) CloseConn() error {
	err := s.DB.Close()
	if err != nil {
		return err
	}
	return nil
}
func (s *SqlConn) GetCharacter() []m.Character {
	results, err := s.DB.Query("SELECT * FROM characters")
	if err != nil {
		fmt.Println("Error", err.Error())
		return nil
	}

	characters := []m.Character{}
	for results.Next() {
		var character m.Character
		err = results.Scan(&character.Name, &character.ID, &character.Class, &character.Race, &character.Level, &character.CreatedAt, &character.Alive)
		if err != nil {
			panic(err.Error())
		}
		characters = append(characters, character)
	}
	return characters
}

// GetCharacters returns an array of characters associated with the account accOwner.
func (s *SqlConn) GetCharacters(code string) []*m.Character {
	results, err := s.DB.Query(fmt.Sprintf("SELECT * FROM characters WHERE characterowner = '%s'", code))
	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}
	fmt.Println()
	// Append every character to the temporary storage
	var charTempStorage []*m.Character
	for {
		char := &m.Character{}
		if results.Next() {
			err = results.Scan(&char.ID, &char.Name, &char.Class, &char.Race, &char.Level, &char.CreatedAt, &char.Alive, &char.CharacterOwner)
			char.CharacterOwner = code
			charTempStorage = append(charTempStorage, char)
			if err != nil {
				return nil
			}
		} else {
			break
		}
	}
	return charTempStorage
}

func (s *SqlConn) AddCharacter(Character m.Character) {
	insert, err := s.DB.Query(
		"INSERT INTO characters (name,id,class,level,race,characterowner) VALUES (?,?,?,?,?,?)",
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
		"DELETE FROM characters WHERE id = ?", Character.ID)
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

//TODO: Convert to postgres
