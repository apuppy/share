package middlewares

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CORSConfig() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowCredentials = true
	config.AllowHeaders = []string{"Content-Type", "Content-Length", "Accept-Encoding", "Accept", "Origin", "Cache-Control", "Authorization", "Access-Control-Allow-Origin"}
	config.AllowMethods = []string{"OPTIONS", "GET", "POST", "HEAD", "PUT", "DELETE"}
	return cors.New(config)
}
