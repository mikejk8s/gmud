package charactersroutes

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"errors"
	m "github.com/mikejk8s/gmud/pkg/models"
	db "github.com/mikejk8s/gmud/pkg/mysql"
	)


 var Characters = []m.Character{
// 	{ID: "1", Name: "John Doe", Class: "Warrior", Race: "Human", Level: 1},
}

func GetCharacters(c *gin.Context) {
	name := c.Param("name")
	for _, Character := range Characters {
		if Character.Name == name {
	c.JSON(http.StatusOK, Characters)
	db.GetCharacters(Character.Name)
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "Character Names not found"})
}


func GetCharacter(c *gin.Context) {
	id := c.Param("id")
	for _, Character := range Characters {
		if Character.ID == id {
			c.JSON(http.StatusOK, Character)
			//return
			db.GetCharacters(id)
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Character ID not found"})
}

func CreateCharacter(c *gin.Context) {
	var Character m.Character
	c.BindJSON(&Character)
	Characters = append(Characters, Character)
	c.JSON(http.StatusCreated, Character)
	db.AddCharacter(Character)
}

func UpdateCharacters(c *gin.Context) {
	id := c.Param("id")
	var Character m.Character
	c.BindJSON(&Character)
	for index, item := range Characters {
		if item.ID == id {
			Characters[index] = Character
			c.JSON(http.StatusOK, Character)
			//return
		}
	}
	c.JSON(http.StatusNotFound, errors.New("Character not found"))
}

func DeleteCharacter(c *gin.Context) {
	var Character m.Character
	id := c.Param("id")
	for index, item := range Characters {
		if item.ID == id {
			Characters = append(Characters[:index], Characters[index+1:]...)
			c.JSON(http.StatusOK, gin.H{"success": "Character deleted"})
			//return
			db.DeleteCharacter(Character)
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
