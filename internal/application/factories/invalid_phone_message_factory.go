package factories

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type InvalidPhoneMessageFactory struct {
}

func (factory InvalidPhoneMessageFactory) CreateMessage(chatId int64) *tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(chatId, "Номер телефона не соответствует указанному формату, попробуйте еще раз")

	replyKeyboard := tgbotapi.NewRemoveKeyboard(true)
	msg.ReplyMarkup = replyKeyboard

	return &msg
}
