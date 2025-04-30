package service

import (
	"log"
	"github.com/Youknow2509/discord_webhook/internal/model"
)

// interface for sending messages with webhook discord
type ISendMessage interface {
	SendText(message model.MessageText) error
	// TODO
}

var vISendMessage ISendMessage

// get instance of ISendMessage
func GetSendMessage() ISendMessage {
	if vISendMessage == nil {
		log.Fatalln("ISendMessage is nil")
	}
	return vISendMessage
}

// set instance of ISendMessage
func SetSendMessage(i ISendMessage) {
	if vISendMessage != nil {
		log.Fatalln("ISendMessage is already set")
	}
	vISendMessage = i
}