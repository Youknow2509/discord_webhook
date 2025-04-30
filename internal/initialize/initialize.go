package initialize

import "github.com/gin-gonic/gin"

func Run() *gin.Engine{
	// Get config
	GetConfig()

	// Initialize service
	InitializeService()

	// Initialize router
	engine := InitializeRouter()

	return engine
}