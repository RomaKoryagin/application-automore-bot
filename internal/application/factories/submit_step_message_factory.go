package factories

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type SubmitStepMessageFactory struct {
}

func (factory SubmitStepMessageFactory) CreateMessage(chatId int64) *tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(chatId, "Подтвердите создание заявки")

	submitApplicationKeyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Оставить контакт"),
		),
	)

	msg.ReplyMarkup = submitApplicationKeyboard

	return &msg
}
