package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mikejk8s/gmud/middlewares"
	"github.com/mikejk8s/gmud/controllers"
	"github.com/mikejk8s/gmud/pkg/userdb"
	cr "github.com/mikejk8s/gmud/pkg/charactersroutes"
)

func ConnectUserDB() {
	// Initialize Database
	userdb.Connect("user:password@tcp(localhost:3307)/users?parseTime=true")
	userdb.Migrate()
	r := InitRouter()
	r.Run(":8080")
}

func InitRouter() *gin.Engine {
	r := gin.Default()
	a := r.Group("/api")
	{
		a.POST("/token", controllers.GenerateToken)
		a.POST("/user/register", controllers.RegisterUser)
		r.GET("/characters", cr.GetCharacters)
		s := a.Group("/secured").Use(middlewares.Auth())
		{
			s.GET("/user", controllers.GetUser)
			s.POST("/token", controllers.GenerateToken)
			s.GET("/characters/:id", cr.GetCharacter)
			s.POST("/characters", cr.CreateCharacter)
			s.PUT("/characters/:id", cr.UpdateCharacters)
			s.DELETE("/characters/:id", cr.DeleteCharacter)
		}
	return r
}
}