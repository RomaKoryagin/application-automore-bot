package factories

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type PersonPhoneStepMessageFactory struct {
}

func (factory PersonPhoneStepMessageFactory) CreateMessage(chatId int64) *tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(chatId, "Пожалуйста, укажите Ваш номер телефона в формате 7ХХХХХХХХХ ⚠️Номер нужен для авторизации, проверка защиты от ботов, благодарим за понимание.")

	replyKeyboard := tgbotapi.NewRemoveKeyboard(true)
	msg.ReplyMarkup = replyKeyboard

	return &msg
}
