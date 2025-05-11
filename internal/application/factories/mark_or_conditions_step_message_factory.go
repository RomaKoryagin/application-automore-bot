package factories

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type MarkOfConditionsStepMessageFactory struct {
}

func (factory MarkOfConditionsStepMessageFactory) CreateMessage(chatId int64) *tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(chatId, "Укажите марку и модель автомобиля, если уже знаете, какой Вам нужен. Или напишите Ваши пожелания, требования к автомобилю или цели его эксплуатации")

	replyKeyboard := tgbotapi.NewRemoveKeyboard(true)
	msg.ReplyMarkup = replyKeyboard

	return &msg
}
