package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"zootechs.com/blog/controllers"
)

func main() {
	r := gin.Default()
	r.Use(useCORS())

	r.GET("/ping", controllers.Pong)
	r.GET("/db", controllers.DbDemo)
	r.POST("/link-category", controllers.SaveCategory)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func useCORS() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "User-Agent", "Referrer", "Host", "Token"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
}
