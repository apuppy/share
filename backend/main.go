package main

import (
	"github.com/gin-gonic/gin"
	"zootechs.com/blog/controllers"
	"zootechs.com/blog/middlewares"
)

func main() {
	r := gin.Default()
	r.Use(middlewares.CORSConfig())

	r.GET("/ping", controllers.Pong)
	r.GET("/db", controllers.DbDemo)
	r.POST("/link-category", controllers.SaveCategory)
	r.GET("/zap-log", controllers.ZapLog)
	r.POST("/editor/save", controllers.SaveContent)
	r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
