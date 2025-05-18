package factories

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type AboutMessageFactory struct {
}

func (factory AboutMessageFactory) CreateMessage(chatId int64) *tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(chatId, "Два адреса:\n\t1. г. Владивосток, ул. Коммунаров, 21 \n\t2. г. Артем, ул. Урбанского, 9")

	replyKeyboard := tgbotapi.NewRemoveKeyboard(true)
	msg.ReplyMarkup = replyKeyboard

	return &msg
}
