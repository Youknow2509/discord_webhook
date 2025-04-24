package impl

import (
	"bytes"
	"encoding/json"
	"log"

	"github.com/Youknow2509/discord_webhook/internal/model"
	"github.com/Youknow2509/discord_webhook/internal/service"
	executewebhook "github.com/Youknow2509/discord_webhook/internal/utils/executeWebhook"
)

// send message impl type
type SendMessageImpl struct {
	UrlWebhook string
}

// SendText implements service.ISendMessage.
func (s *SendMessageImpl) SendText(message string) error {
	messageSend := model.MessageText{
		UserName: "Test user",
		Content: message,
		AvatarUrl: "https://cdn.discordapp.com/avatars/123456789012345678/abcdef1234567890abcdef1234567890.png",
	}
	body := new(bytes.Buffer)
	err := json.NewEncoder(body).Encode(messageSend)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(body)

	res, err := executewebhook.ExcuteWebhook(s.UrlWebhook, body)
	if err != nil {
		log.Fatal("client: error making http request: %s\n", err)
	}
	if res.StatusCode != 204 {
		log.Printf("client: non 204 response: %s\n", res.Status)
	}
	defer res.Body.Close()
	return nil
}

// new instance of SendMessageImpl
func NewSendMessageImpl(urlWebhook string) service.ISendMessage {
	return &SendMessageImpl{
		UrlWebhook: urlWebhook,
	}
}
