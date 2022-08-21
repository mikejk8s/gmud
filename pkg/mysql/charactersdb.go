package mysql

import (
	"database/sql"
	//_ "github.com/go-sql-driver/mysql"
	//"github.com/mikejk8s/gmud"
	"fmt"

	//cr "github.com/mikejk8s/gmud/pkg/charactersroutes"
	m "github.com/mikejk8s/gmud/pkg/models"
)

func GetCharacters() []m.Character {
	db, err := sql.Open("mysql", dbuser+":"+dbpass+"@tcp(127.0.0.1:3306)/"+dbname)
	if err != nil {
		fmt.Println("Error", err.Error())
		return nil
}

defer db.Close()
results, err := db.Query("SELECT * FROM character")
	if err != nil {
		fmt.Println("Error", err.Error())
		return nil
		}

		characters := []m.Character{}
		for results.Next() {
			var character m.Character
			err = results.Scan(&character.Name, &character.ID, &character.Class, &character.Level)
			if err!=nil {
				panic(err.Error())
				}
				characters = append(characters, character)
				}
				return characters
				}

				func GetCharacter(code string) *m.Character {
					db, err := sql.Open("mysql", dbuser+":"+dbpass+"@tcp(127.0.0.1:3306)/"+dbname)
					char:= &m.Character{}
					if err != nil {
						// simply print the error to the console
						fmt.Println("Err", err.Error())
						// returns nil on error
						return nil
					}
					defer db.Close()
					results, err := db.Query("SELECT * FROM Character where code=?", code)
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
					db, err := sql.Open("mysql", dbuser+":"+dbpass+"@tcp(127.0.0.1:3306)/"+dbname)
					if err != nil {
						panic(err.Error())
					}
					// defer the close till after this function has finished
					// executing
					defer db.Close()
					insert, err := db.Query(
						"INSERT INTO Character (name,id,class,level) VALUES (?,?,?, now())",
						Character.Name, Character.Class, Character.Level)
					// if there is an error inserting, handle it
					if err != nil {
						panic(err.Error())
					}
					defer insert.Close()
				}

//TODO: Convert to postgres