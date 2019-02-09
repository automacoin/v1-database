package api

import (
	"github.com/gin-gonic/gin"
)

// setRoutes associates routes and handlers
func setRoutes(router *gin.Engine) {
	// TODO
	// some routes will be enabled/disabled on environment

	router.GET("/ping", pingHandler)

	return
}

// pingHandler returns "pong"
func pingHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
