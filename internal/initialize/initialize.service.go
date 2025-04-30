package initialize

import (
	"github.com/Youknow2509/discord_webhook/internal/global"
	"github.com/Youknow2509/discord_webhook/internal/service"
	"github.com/Youknow2509/discord_webhook/internal/service/impl"
)

// Initialize service
func InitializeService() {
	// Initialize impl service send message
	implSendMessage := impl.NewSendMessageImpl(global.URL_WEBHOOK)
	// initialize interface send message
	service.SetSendMessage(implSendMessage)
}