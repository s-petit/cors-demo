package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	router := gin.Default()

	router.LoadHTMLGlob("html/*")

	router.GET("/malevolent", func(c *gin.Context) {
		c.HTML(http.StatusOK, "malevolent.html", gin.H{})
	})

	router.Run()

}