package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Success success http response
func Success(c *gin.Context, data interface{}) {
	code := 20000
	Response(c, code, "", data)
}

// Error error http response
func Error(c *gin.Context, message string) {
	code := 40000
	Response(c, code, message, nil)
}

// Response HTTP response
func Response(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(
		http.StatusOK,
		gin.H{"code": code, "message": message, "data": data},
	)
}
