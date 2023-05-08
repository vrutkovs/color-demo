package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/penglongli/gin-metrics/ginmetrics"
)

func main() {
	router := gin.Default()

	router.ForwardedByClientIP = true

	router.Use(GinContextToContextMiddleware())

	loadAssets(router)
	exposeMetrics(router)

	router.GET("/", Home)
	router.GET("/healthz", healthz)

	router.Run()

}

func loadAssets(router *gin.Engine) {
	base_path := os.Getenv("BASE_PATH")
	if len(base_path) == 0 {
		base_path = "./"
	}
	templates_path := filepath.Join(base_path, "templates")
	router.LoadHTMLGlob(fmt.Sprintf("%s/*", templates_path))
}

func exposeMetrics(router *gin.Engine) {
	metrics := ginmetrics.GetMonitor()
	metrics.SetMetricPath("/metrics")
	metrics.SetSlowTime(5)
	metrics.SetDuration([]float64{0.1, 0.3, 1.2, 5, 10})
	metrics.Use(router)
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
