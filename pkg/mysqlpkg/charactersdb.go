package mysqlpkg

import (
	"database/sql"
	"log"

	//_ "github.com/go-sql-driver/mysql"
	//"github.com/mikejk8s/gmud"
	"fmt"

	"github.com/felixge/fgtrace"
	//cr "github.com/mikejk8s/gmud/pkg/charactersroutes"
	m "github.com/mikejk8s/gmud/pkg/models"
)

// func GetCharacters(code string) []m.Character {
// 	db, err := sql.Open("mysql", username+":"+password+"@tcp(127.0.0.1:3306)/"+dbname)
// 	char:= &m.Character{}
// 	if err != nil {
// 		fmt.Println("Error", err.Error())
// 		return nil
// 		{
// 	defer db.Close()
// 	results, err := db.Query("SELECT * FROM characters")
// 	if err != nil {
// 		fmt.Println("Err", err.Error())
// 		return nil
// 	}
// 	defer results.Close()
// 	for result.Next()

func GetCharacter() []m.Character {
	db, err := sql.Open("mysql", username+":"+password+"@tcp(127.0.0.1:3306)/"+dbname+"?parseTime=true")
	if err != nil {
		fmt.Println("Error", err.Error())
		return nil
	}

	defer func(trace *fgtrace.Trace) {
		err := trace.Stop()
		if err != nil {
			log.Println(err)
		}
	}(fgtrace.Config{Dst: fgtrace.File("charactersdb-fgtrace.json")}.Trace())

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Println(err)
		}
	}(db)
	results, err := db.Query("SELECT * FROM characters")
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

func GetCharacters(code string) *m.Character {
	db, err := sql.Open("mysql", username+":"+password+"@tcp(127.0.0.1:3306)/"+dbname+"?parseTime=true")
	char := &m.Character{}
	if err != nil {
		log.Println("Error", err.Error())
	}
	defer db.Close()
	results, err := db.Query(fmt.Sprintf("SELECT * FROM characters WHERE characterowner = '%s'", code))
	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}
	if results.Next() {
		err = results.Scan(&char.ID, &char.Name, &char.Class, &char.Race, &char.Level, &char.CreatedAt, &char.Alive, &char.CharacterOwner)
		char.CharacterOwner = code
		if err != nil {
			return nil
		}
	} else {
		return nil
	}
	return char
}

func AddCharacter(Character m.Character) {
	db, err := sql.Open("mysql", username+":"+password+"@tcp(127.0.0.1:3306)/"+dbname)
	if err != nil {
		panic(err.Error())
	}
	// defer the close till after this function has finished
	// executing
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Println(err)
		}
	}(db)
	insert, err := db.Query(
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

func DeleteCharacter(Character m.Character) {
	db, err := sql.Open("mysql", username+":"+password+"@tcp(127.0.0.1:3306)/"+dbname)
	if err != nil {
		panic(err.Error())
	}
	// defer the close till after this function has finished
	// executing
	defer db.Close()
	delete, err := db.Query(
		"DELETE FROM characters WHERE id = ?", Character.ID)
	// if there is an error deleting, handle it
	if err != nil {
		panic(err.Error())
	}
	defer delete.Close()
}

//TODO: Convert to postgres
