package initialize

import (
	"github.com/gin-gonic/gin"
	routerApp "github.com/Youknow2509/discord_webhook/internal/router"
)

// initialize router
func InitializeRouter() *gin.Engine{
	router := gin.Default()

	// set global middleware
	router.Use()

	//
	r := routerApp.RouterGroupApp

	// set the router group
	MainRouterGroup := router.Group("/api/")
	{
		r.SendMessage.SendMessageRouter(MainRouterGroup)
	}

	return router
}