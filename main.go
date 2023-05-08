package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.ForwardedByClientIP = true

	router.Use(GinContextToContextMiddleware())
	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("templates/*")

	router.GET("/", Home)

	router.Run()

}

func GinContextToContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("color", os.Getenv("COLOR"))
		c.Set("title", os.Getenv("TITLE"))
		c.Next()
	}
}

func Home(c *gin.Context) {
	title, ok := c.Get("title")
	if !ok || title == "" {
		title = "ERROR UNKNOWN TITLE"
	}

	color, ok := c.Get("color")
	if !ok || color == "" {
		color = "red"
	}

	c.HTML(http.StatusOK, "home.html", gin.H{
		"title": title,
		"color": color,
	})
}
