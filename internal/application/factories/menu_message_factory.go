package factories

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type MenuMessageFactory struct {
}

func (factory MenuMessageFactory) CreateMessage(chatId int64) *tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(chatId, "Вы в меню! Пожалуйста, выберите интересующий вас пункт меню.")

	countriesKeyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Сайт", "/link"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Как нас найти", "/about"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Оставить заявку", "/newapplication"),
		),
	)

	msg.ReplyMarkup = countriesKeyboard

	return &msg
}
