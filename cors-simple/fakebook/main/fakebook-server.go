package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/gin-contrib/cors"
)

func main() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
	}))

	router.LoadHTMLGlob("html/*")

	router.GET("/fakebook", func(c *gin.Context) {
		c.HTML(http.StatusOK, "fakebook.html", gin.H{"name": "Jack Lee"})
	})

	router.Run(":8081")
}