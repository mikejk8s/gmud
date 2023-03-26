package backend

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
	"github.com/mikejk8s/gmud/pkg/postgrespkg"
	"golang.org/x/crypto/bcrypt"
)

// SignupFormJSONBinding sets JSON data that has arrived from signup.html's fetch request.
//
// Check cmd/app/templates for the signup.html.
func SignupFormJSONBinding(c *gin.Context) {
	var loginReq = new(postgrespkg.LoginReq)
	if err := c.BindJSON(&loginReq); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed"})
		return
	}
	if err := hashAndSalt([]byte(loginReq.Password), 16, loginReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed"})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "success"})
	}
}
func hashAndSalt(pwd []byte, minCost int, userInfo *postgrespkg.LoginReq) error {
	hash, err := bcrypt.GenerateFromPassword(pwd, minCost)
	if err != nil {
		log.Println(err)
		return err
	}

	dbUsers := postgrespkg.SqlConn{}
	if err := dbUsers.GetSQLConn("users"); err != nil {
		log.Println(err)
		return err
	}
	defer dbUsers.Close()

	userInfo.Password = string(hash)

	if err := dbUsers.CreateNewUser(*userInfo); err != nil {
		return err
	}

	return nil
}

// SignupPage is a literally fleshed out signup page just consisting three input fields with a submit button.
func SignupPage(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"signup.html",
		gin.H{
			"CallbackURL": "callback",
		},
	)
}

// StartWebPageBackend gets WEBPAGE_HOST and WEBPAGE_PORT and listens a new HTTP server from these values.
//
// You can set these values in the docker-compose.yml file.
//
// If app is not ran on Docker (envExists=false), it will use the localPort variable that is passed to the function.
func StartWebPageBackend(envExists bool, localPort int) {
	r := gin.Default()
	// This is used for hiding printing one hundred of lines of loading static files.
	// If you want to see which files are loaded you can remove this line.
	gin.SetMode(gin.ReleaseMode)
	r.GET("/signup", SignupPage)
	r.POST("/callback", SignupFormJSONBinding)
	r.GET("/exists", func(c *gin.Context) {
		c.HTML(
			http.StatusOK,
			"exists.html",
			gin.H{},
		)
	})
	r.GET("/success", func(c *gin.Context) {
		c.Writer.Write([]byte("Success!"))
	})
	r.HTMLRender = ginview.Default()
	if envExists {
		// If we are running from Docker, then working directory consists all folders.
		// Means, cmd/app/templates is the correct path.
		r.LoadHTMLGlob("cmd/app/templates/*.html")
		r.Run(fmt.Sprintf("%s:%s", os.Getenv("WEBPAGE_HOST"), os.Getenv("WEBPAGE_PORT")))
	} else {
		r.LoadHTMLGlob("templates/*.html")
		r.Run(fmt.Sprintf(":%d", localPort))
	}
	// r.Run(":6969")
}
