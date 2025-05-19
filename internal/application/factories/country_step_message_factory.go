package factories

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type CountryStepMessageFactory struct {
}

func (factory *CountryStepMessageFactory) CreateMessage(chatId int64) *tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(chatId, "Приветствуем Вас в компании «AUTOMORE»  (Авто Море) ! Автомобили ведущих мировых автопроизводителей для Вас «под ключ»! Подберём. Проверим. Доставим. Оформим. \n\nВыберите из списка страну, где планируете приобрести автомобиль:")

	countriesKeyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Япония", "/japan"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Корея", "/korea"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Китай", "/china"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("ГЛАВНОЕ МЕНЮ", "/menu"),
		),
	)

	msg.ReplyMarkup = countriesKeyboard

	return &msg
}

func NewCountryStepMessageFactory() *CountryStepMessageFactory {
	return &CountryStepMessageFactory{}
}
