package controllers

import (
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func SaveContent(c *gin.Context) {
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		panic(err)
	}
	demoResponse := map[string]string{
		"url":         c.Request.URL.String(),
		"get_params":  c.Request.URL.Query().Encode(),
		"post_params": string(jsonData),
	}
	Success(c, demoResponse)
}
