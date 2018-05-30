package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/sessions"
	"golang.org/x/crypto/bcrypt"
)

const appLogin = "user"
const appPassword = "$2a$08$QG4OYI7WpXKaTwUMiD2/XeNLZAycV14bl/j4F5gMQzOc9XdK7AgvO;" // pass value is: lol


func main() {
	router := gin.Default()

	store := sessions.NewCookieStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	router.POST("/login", func(c *gin.Context) {

		login := c.PostForm("login")
		pass := c.PostForm("pass")

		session := sessions.Default(c);
		session.Set("login", login)
		session.Set("pass", pass)

		session.Save()

		c.Redirect(http.StatusFound, "/fakebook")
	})

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{})
	})

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://clickbait:8080"},
		AllowMethods: []string{"GET", "POST"},
		AllowCredentials: true,
	}))

	router.LoadHTMLGlob("html/*")

	router.GET("/fakebook", func(c *gin.Context) {

		session := sessions.Default(c)
		sessionLogin, sessionPass := getCredentialsFromSession(session)
		authorized := isAuthorized(sessionLogin, sessionPass, c)

		if (authorized) {
			c.HTML(http.StatusOK, "fakebook.html", gin.H{"name": "Jack Lee"})
		} else {
			c.String(http.StatusUnauthorized, "unauthorized")
		}
		//c.SetCookie("malevolent", hack, 100, "/", "", false, true)
	})

	//TODO : preflight , and withCredentials and cookie attack
	// CORS et proxy : pourquoi cela regle le probleme

	router.Run(":8081")
}

func getCredentialsFromSession(session sessions.Session) (string, string) {
	sessionLogin := session.Get("login")
	if sessionLogin == nil {
		sessionLogin = ""
	}
	sessionPass := session.Get("pass")
	if sessionPass == nil {
		sessionPass = ""
	}

	return sessionLogin.(string), sessionPass.(string)
}

func isAuthorized(login string, pass string, c *gin.Context) (bool) {
	// Compare the stored hashed password, with the hashed version of the password that was received
	err := bcrypt.CompareHashAndPassword([]byte(appPassword), []byte(pass))

	return err == nil && login == appLogin
}