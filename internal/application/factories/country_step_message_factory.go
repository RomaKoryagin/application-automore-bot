package factories

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type CountryStepMessageFactory struct {
}

func (factory CountryStepMessageFactory) CreateMessage(chatId int64) *tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(chatId, "Приветствует Вас ! Спасибо за обращение в компанию «AUTOMORE»  (Автоморе) !Мы занимаемся доставкой автомобилей ведущих мировых автопроизводителей, включая европейских, и стран Азии.Подберём. Проверим. Доставим. Оформим. Как для себя!Пожалуйста, выберите из списка страну, где хотите приобрести автомобиль:")

	countriesKeyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Япония"),
			tgbotapi.NewKeyboardButton("Корея"),
			tgbotapi.NewKeyboardButton("Китай"),
		),
	)

	msg.ReplyMarkup = countriesKeyboard

	return &msg
}
