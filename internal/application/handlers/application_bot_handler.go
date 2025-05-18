package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type IApplicationBotHandlerTgMsgService interface {
	SendReplyMessage(chatId int64, telegramId string, text string)
}

type ApplicationBotHandler struct {
	TgMsgService IApplicationBotHandlerTgMsgService
}

func (handler *ApplicationBotHandler) Handle(c *gin.Context) {
	var update tgbotapi.Update

	if err := c.BindJSON(&update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var msg string
	var chatId int64
	var username string

	if update.CallbackQuery != nil {
		chatId = update.CallbackQuery.From.ID
		msg = update.CallbackQuery.Data
		username = update.CallbackQuery.From.UserName
	}

	if update.Message != nil {
		msg = update.Message.Text
		chatId = update.Message.Chat.ID
		username = update.Message.From.UserName
	}

	go handler.TgMsgService.SendReplyMessage(chatId, username, msg)
}

func NewApplicationBotHandler(
	tgMsgService IApplicationBotHandlerTgMsgService,
) *ApplicationBotHandler {
	return &ApplicationBotHandler{TgMsgService: tgMsgService}
}
