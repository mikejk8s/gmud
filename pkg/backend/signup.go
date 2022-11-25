package backend

import (
	"fmt"
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
	"github.com/mikejk8s/gmud/pkg/mysqlpkg"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"os"
)

// SignupFormJSONBinding sets JSON data that has arrived from signup.html's fetch request.
//
// Check cmd/app/templates for the signup.html.
func SignupFormJSONBinding(c *gin.Context) {
	// type LoginReq struct {
	//		Name     string `json:"username"`
	//		Password string `json:"password"`
	//		Email    string `json:"email"`
	//	}
	var LoginJSON = mysqlpkg.LoginReq{}
	// Bind the josn to the user credentials struct.
	err := c.BindJSON(&LoginJSON)
	if err != nil {
		fmt.Println(err)
	}
	// Hash the password and salt it with 16 min cost, this can change. Then create a new user with the LoginJSON struct.
	hashAndSalt([]byte(LoginJSON.Password), 16, LoginJSON)
}
func hashAndSalt(pwd []byte, minCost int, userInfo mysqlpkg.LoginReq) {

	// Use GenerateFromPassword to hash & salt pwd.
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided it isn't lower
	// than the MinCost (4)
	hash, err := bcrypt.GenerateFromPassword(pwd, minCost)
	if err != nil {
		log.Println(err)
	}
	// GenerateFromPassword returns a byte slice, so we need to
	// convert the bytes to a string and return it

	// get a new db connection to the user table
	dbUsers := mysqlpkg.SqlConn{}
	err = dbUsers.GetSQLConn("users")
	if err != nil {
		log.Println(err)
	}
	// Set the user credential structs password to the new hashed password.
	userInfo.Password = string(hash)
	// Create a new user from all these information.
	err = dbUsers.CreateNewUser(userInfo)
	if err != nil {
		log.Println(err)
	}
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
// Or just comment it out and use r.Run(":desiredport") to run the server on a desired port.
func StartWebPageBackend() {
	r := gin.Default()
	r.GET("/signup", SignupPage)
	r.POST("/callback", SignupFormJSONBinding)
	r.HTMLRender = ginview.Default()
	r.LoadHTMLGlob("templates/*.html")
	r.Run(fmt.Sprintf("%s:%s", os.Getenv("WEBPAGE_HOST"), os.Getenv("WEBPAGE_PORT")))
	// r.Run(":6969")
}
