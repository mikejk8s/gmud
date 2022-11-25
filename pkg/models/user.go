package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Username string `json:"username" gorm:"unique"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}

// CheckPassword gets plain password as input and checks if it matches the hashed password in the database.
//
// user.Password is set during fetching the users database, and retrieved as already hashed.
//
// If err is not nil, then the password is not correct and SSH password authentication will fail by returning false.
func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
