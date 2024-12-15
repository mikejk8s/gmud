package routes

import (
	"net/http"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/semihalev/gin-stats"
	nrgin "github.com/newrelic/go-agent/v3/integrations/nrgin"
	"github.com/newrelic/go-agent/v3/newrelic"

	"github.com/mikejk8s/gmud/controllers"
	"github.com/mikejk8s/gmud/middlewares"
	cr "github.com/mikejk8s/gmud/pkg/charactersroutes"
	"github.com/mikejk8s/gmud/pkg/userdb"
)

// TODO: Hardcoded db endpoints
func ConnectUserDB() {
	// Initialize Database
	userdb.Connect("user:password@tcp(localhost:3307)/users?parseTime=true")
	userdb.Migrate()
	r := InitRouter()
	r.Run(":8080")
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