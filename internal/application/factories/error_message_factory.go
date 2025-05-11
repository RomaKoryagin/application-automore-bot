package factories

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type ErrorMessageFactory struct {
}

func (factory ErrorMessageFactory) CreateMessage(chatId int64) *tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(chatId, "Произошла непредвиденная ошибка... Постараемся исправить ее в ближайшее время")

	replyKeyboard := tgbotapi.NewRemoveKeyboard(true)
	msg.ReplyMarkup = replyKeyboard

	return &msg
}

func NewErrorMessageFactory() *ErrorMessageFactory {
	return &ErrorMessageFactory{}
}
