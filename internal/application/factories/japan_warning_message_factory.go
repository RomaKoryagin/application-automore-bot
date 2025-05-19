package factories

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type JapanWarningMessageFactory struct {
}

func (factory JapanWarningMessageFactory) CreateMessage(chatId int64) *tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(chatId, "У всех АВТО в Японии, рулевая колонка расположена справа. Поэтому, если вы хотите японский авто, но с рулевой колонкой, расположенной слева, то пройдите в раздел Корея или Китай.")

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
