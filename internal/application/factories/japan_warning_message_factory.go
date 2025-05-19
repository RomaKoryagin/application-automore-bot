package factories

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type JapanWarningMessageFactory struct {
}

func (factory JapanWarningMessageFactory) CreateMessage(chatId int64) *tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(chatId, "У всех автомобилей из Японии, рулевая колонка расположена справа.\nЕсли вы хотите автомобиль японской марки, но с рулевой колонкой слева, перейдите в раздел Корея или Китай")

	replyKeyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Вернуться к выбору страны", "/return-country-step"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Расположение руля справа", "/submit-right-wheeling-type"),
		),
	)

	msg.ReplyMarkup = replyKeyboard

	return &msg
}

func NewJapanWarningMessageFactory() *JapanWarningMessageFactory {
	return &JapanWarningMessageFactory{}
}
