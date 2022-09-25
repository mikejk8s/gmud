// https://github.com/iammukeshm/jwt-authentication-golang

package userdb //TODO: Refactor this and mysql package to not collide

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mikejk8s/gmud/middlewares"
    "github.com/mikejk8s/gmud/controllers"
	"github.com/mikejk8s/gmud/pkg/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Instance *gorm.DB
var dbError error

func Connect(connectionString string) () {
	Instance, dbError = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if dbError != nil {
		log.Fatal(dbError)
		panic("Cannot connect to DB")
	}
	log.Println("Connected to Database!")
}

func Migrate() {
	Instance.AutoMigrate(&models.User{})
	log.Println("Database Migration Completed!")
}

func ConnectUserDB() {
	// Initialize Database
	Connect("user:password@tcp(localhost:3307)/users?parseTime=true")
	Migrate()
	router := initRouter()
	router.Run(":8080")
}

func initRouter() *gin.Engine {
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

