package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"errors"
	)


type Character struct {
	Name string	`json:"name`
	ID string	`json:"id"`
	Class string	`json:"class"`
	Race string	`json:"race"`
	Level int	`json:"level"`
}

var Characters = []Character{
	{ID: "1", Name: "John Doe", Class: "Warrior", Race: "Human", Level: 1},
	{ID: "2", Name: "Jacbo Woo", Class: "Wizard", Race: "Human", Level: 2},
	{ID: "3", Name: "Carly Howe", Class: "Warrior", Race: "Human", Level: 5},
}

func GetCharacters(c *gin.Context) {
	c.JSON(http.StatusOK, Characters)
}

func GetCharacter(c *gin.Context) {
	id := c.Param("id")
	for _, Character := range Characters {
		if Character.ID == id {
			c.JSON(http.StatusOK, Character)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Character not found"})
}

func CreateCharacter(c *gin.Context) {
	var Character Character
	c.BindJSON(&Character)
	Characters = append(Characters, Character)
	c.JSON(http.StatusCreated, Character)
}

func UpdateCharacters(c *gin.Context) {
	id := c.Param("id")
	var Character Character
	c.BindJSON(&Character)
	for index, item := range Characters {
		if item.ID == id {
			Characters[index] = Character
			c.JSON(http.StatusOK, Character)
			return
		}
	}
	c.JSON(http.StatusNotFound, errors.New("Character not found"))
}

func DeleteCharacter(c *gin.Context) {
	id := c.Param("id")
	for index, item := range Characters {
		if item.ID == id {
			Characters = append(Characters[:index], Characters[index+1:]...)
			c.JSON(http.StatusOK, gin.H{"success": "Character deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, errors.New("Character not found"))
}

func CharactersRoutes() {
	r := gin.Default()
	r.GET("/characters", GetCharacters)
	r.GET("/characters/:id", GetCharacter)
	r.POST("/characters", CreateCharacter)
	r.PUT("/characters/:id", UpdateCharacters)
	r.DELETE("/characters/:id", DeleteCharacter)
	r.Run(":8080")
}
