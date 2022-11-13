package routes

import (
	"context"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/semihalev/gin-stats"

	"github.com/mikejk8s/gmud/controllers"
	"github.com/mikejk8s/gmud/middlewares"
	cr "github.com/mikejk8s/gmud/pkg/charactersroutes"
	"github.com/mikejk8s/gmud/pkg/userdb"

	nrgin "github.com/newrelic/go-agent/v3/integrations/nrgin"
	"github.com/newrelic/go-agent/v3/newrelic"

	"fmt"
	"os"
)

func ConnectUserDB() {
	// Initialize Database
	characterDB, err := gorm.Open(mysql.Open("root:1234@tcp(localhost:3306)/"), &gorm.Config{})
	_, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	characterDB.Exec("CREATE SCHEMA IF NOT EXISTS " + "users")
	if err != nil {
		log.Printf("Error %s when creating DB\n", err)
		panic(err)
	}
	userdb.Connect("root:1234@tcp(localhost:3306)/users?parseTime=true")
	userdb.Migrate()
	r := InitRouter()
	err = r.Run(":8080")
	if err != nil {
		panic(err)
	}
}

func InitRouter() *gin.Engine {

	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName("Gmud"),
		newrelic.ConfigLicense(os.Getenv("NEW_RELIC_API_KEY")),
		newrelic.ConfigDebugLogger(os.Stdout),
		newrelic.ConfigCodeLevelMetricsEnabled(true),
	)
	if nil != err {
		fmt.Println(err)
		os.Exit(1)
	}

	r := gin.Default()
	r.Use(nrgin.Middleware(app))

	// stats / 200 OK
	r.Use(stats.RequestStats())

	r.GET("/stats", func(c *gin.Context) {
		c.JSON(http.StatusOK, stats.Report())
	})

	// Routes
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
