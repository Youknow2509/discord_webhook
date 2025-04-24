package main

import (
	"github.com/Youknow2509/discord_webhook/internal/service"
	"github.com/Youknow2509/discord_webhook/internal/service/impl"
)

func main() {
	urlWebhook := ""
	content := "Hello world"
	sendMessage := impl.NewSendMessageImpl(urlWebhook)
	service.SetSendMessage(sendMessage)
	intance := service.GetSendMessage()

	err := intance.SendText(content)
	if err != nil {
		panic(err)
	}
}
