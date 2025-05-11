package factories

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type NoActiveApplicationMessageFactory struct {
}

func (factory NoActiveApplicationMessageFactory) CreateMessage(chatId int64) *tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(chatId, "Заявка для заполнения данных не найдена, вы можете создать новую, используя команду /newapplication")

	replyKeyboard := tgbotapi.NewRemoveKeyboard(true)
	msg.ReplyMarkup = replyKeyboard

	return &msg
}

func NewNoActiveApplicationMessageFactory() *NoActiveApplicationMessageFactory {
	return &NoActiveApplicationMessageFactory{}
}
