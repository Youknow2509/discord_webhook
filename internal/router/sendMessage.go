package router

import (
	"github.com/Youknow2509/discord_webhook/internal/controller"
	"github.com/gin-gonic/gin"
)

type SendMessageRouterGroup struct {

}

// func router
func (s *SendMessageRouterGroup) SendMessageRouter(Router *gin.RouterGroup) {
	router := Router.Group("v1") 
	{
		router.POST("send_text", controller.SendText)
		// TODO: add more router
	}
}