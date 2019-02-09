package api

import (
	"github.com/gin-gonic/gin"
)

// Setup prepares middleware and routes of the API
func Setup() *gin.Engine {
	// TODO
	// set a flag to switch debug/release in gin
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	// set middlewares
	router.Use(getLogger())
	router.Use(gin.Recovery())
	// TODO
	// authentication middleware

	// set routes
	setRoutes(router)

	// we are good to go
	return router
}

// Run starts the HTTP server
func Run(router *gin.Engine) {
	router.Run()
}
