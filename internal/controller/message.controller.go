package controller

import (
	"net/http"

	"github.com/Youknow2509/discord_webhook/internal/global"
	"github.com/Youknow2509/discord_webhook/internal/model"
	"github.com/Youknow2509/discord_webhook/internal/model/response"
	"github.com/Youknow2509/discord_webhook/internal/service"
	"github.com/gin-gonic/gin"
)

// Controller send text
//	@Summary		Send text to discord with webhook
//	@Description	Send text to discord with webhook
//	@Accept			json
//	@Produce		json
//	@Param			body	body		model.MessageText		true	"Message text"
//	@Success		200		{object}	response.ResponseData 		"Success"
//	@Failure		400		{object}	response.ErrResponseData 	"Bad request"
//	@Router			/api/v1/send_text [post]
func SendText(c *gin.Context) {
	var messageText model.MessageText
	if err := c.ShouldBindJSON(&messageText); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if messageText.Content == "" || messageText.Content == "string" {
		response.ErrorResponse(c, http.StatusBadRequest, "Content is required")
		return
	}

	if messageText.AvatarUrl == "" || messageText.AvatarUrl == "string" {
		messageText.AvatarUrl = global.URL_AVATAR
	}
	if messageText.UserName == "" || messageText.UserName == "string" {
		messageText.UserName = global.DEFAULT_USERNAME
	}

	sendMessage := service.GetSendMessage()
	err := sendMessage.SendText(messageText)
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.SuccessResponse(c, http.StatusOK, nil)
}
