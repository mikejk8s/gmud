package charactersdb

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"github.com/mikejk8s/gmud"
)

CREATE TABLE Character (
	`name` varchar(20) NOT NULL DEFAULT 'x',
	`id` varchar(20) NOT NULL DEFAULT NULL,
	`class` varchar(20) NOT NULL DEFAULT NULL,
	`race` varchar(20) NOT NULL DEFAULT NULL,
	`level` int(3) NOT NULL DEFAULT '1'
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

func Get Characters() []Character {

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

		characters := []Character{}
		for results.Next() {
			var character Character
			err = results.Scan(&character.Name, &character.ID, &character.Class, &character.Level)
			if err!=nil {
				panic(err.Error())

				}

				characters = append(characters, character)
				}

				return characters
				}

				func GetCharacter(code string) *Character {

					db, err := sql.Open("mysql", dbuser+":"+dbpass+"@tcp(127.0.0.1:3306)/"+dbname)
					char:= &Character{}
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
						err = results.Scan(&character.Name, &character.ID, &character.Class, &character.Level)
						if err != nil {
							return nil
						}
					} else {

						return nil
					}

					return char
				}

				func AddCharacter(Character Character) {

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