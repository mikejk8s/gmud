package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mikejk8s/gmud/middlewares"
	"github.com/mikejk8s/gmud/controllers"
	"github.com/mikejk8s/gmud/pkg/userdb"
)

func ConnectUserDB() {
	// Initialize Database
	userdb.Connect("user:password@tcp(localhost:3307)/users?parseTime=true")
	userdb.Migrate()
	router := InitRouter()
	router.Run(":8080")
}

func InitRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/user/register", controllers.RegisterUser)
		api.POST("/token", controllers.GenerateToken)
		secured := api.Group("/secured").Use(middlewares.Auth())
		{
			secured.GET("/user", controllers.GetUser)
			secured.POST("/token", controllers.GenerateToken)
		}
	return router
}
}