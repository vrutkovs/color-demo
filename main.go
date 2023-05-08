package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.ForwardedByClientIP = true
	router.Static("/assets", "./assets")

	router.LoadHTMLGlob("templates/*")

	router.GET("/", Home)

	router.Run()

}

func Home(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", gin.H{
		"title": "hello",
		"color": "red",
	})
}
