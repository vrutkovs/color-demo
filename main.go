package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/penglongli/gin-metrics/ginmetrics"
)

func main() {
	router := gin.Default()

	router.ForwardedByClientIP = true

	router.Use(GinContextToContextMiddleware())
	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("templates/*")

	metrics := ginmetrics.GetMonitor()
	metrics.SetSlowTime(5)
	metrics.SetDuration([]float64{0.1, 0.3, 1.2, 5, 10})
	metrics.Use(router)

	router.GET("/", Home)
	router.GET("/healthz", healthz)

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

func healthz(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
