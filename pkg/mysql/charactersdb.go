package mysql

import (
	"database/sql"
	"fmt"

	//_ "github.com/go-sql-driver/mysql"
	"github.com/felixge/fgtrace"

	//"github.com/mikejk8s/gmud"
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

// TODO: db endpoints hardcoded
func GetCharacter() []m.Character {
	db, err := sql.Open("mysql", username+":"+password+"@tcp(127.0.0.1:3306)/"+dbname+"?parseTime=true")
	if err != nil {
		fmt.Println("Error", err.Error())
		return nil
	}

	defer fgtrace.Config{Dst: fgtrace.File("charactersdb-fgtrace.json")}.Trace().Stop()

	defer db.Close()
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
	db, err := sql.Open("mysql", username+":"+password+"@tcp(127.0.0.1:3306)/"+dbname)
	char := &m.Character{}
	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}
	defer db.Close()
	results, err := db.Query("SELECT * FROM characters WHERE id = ?", code)
	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}
	if results.Next() {
		err = results.Scan(&char.Name, &char.ID, &char.Class, &char.Level)
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
	defer db.Close()
	insert, err := db.Query(
		"INSERT INTO characters (name,id,class,level,race) VALUES (?,?,?,?,?)",
		Character.Name, Character.ID, Character.Class, Character.Level, Character.Race)
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
}

func DeleteCharacter(Character m.Character) {
	db, err := sql.Open("mysql", username+":"+password+"@tcp(127.0.0.1:3306)/"+dbname)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	delete, err := db.Query(
		"DELETE FROM characters WHERE id = ?", Character.ID)
	if err != nil {
		panic(err.Error())
	}
	defer delete.Close()
}

//TODO: Convert to postgres TODO: Reconsider this it was less fun than intended
