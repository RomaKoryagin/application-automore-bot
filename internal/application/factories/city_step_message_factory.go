package factories

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type CityStepMessageFactory struct {
}

func (factory CityStepMessageFactory) CreateMessage(chatId int64) *tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(chatId, "Укажите Ваш город, чтобы мы могли учитывать часовой пояс")
	replyKeyboard := tgbotapi.NewRemoveKeyboard(true)
	msg.ReplyMarkup = replyKeyboard

	return &msg
}

func NewCityStepMessageFactory() *CityStepMessageFactory {
	return &CityStepMessageFactory{}
}
