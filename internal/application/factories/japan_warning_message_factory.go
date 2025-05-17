package factories

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type JapanWarningMessageFactory struct {
}

func (factory JapanWarningMessageFactory) CreateMessage(chatId int64) *tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(chatId, "«Расположение руля автомобиля слева или справа? ⚠️ Внимание! ⚠️ У всех АВТО в Японии, рулевая колонка расположена справа. Поэтому, если вы хотите японский авто, но с рулевой колонкой, расположенной слева, то пройдите в раздел Корея или Китай.")

	replyKeyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Вернуться к выбору страны", "Вернуться к выбору страны"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Расположение руля справа", "Расположение руля справа"),
		),
	)

	msg.ReplyMarkup = replyKeyboard

	return &msg
}
