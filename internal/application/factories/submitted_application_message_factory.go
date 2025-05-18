package factories

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type SubmittedApplicationMessageFactory struct {
}

func (factory SubmittedApplicationMessageFactory) CreateMessage(chatId int64) *tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(chatId, "Спасибо ! Ваша заявка принята в работу. В ближайшее время, наш менеджер свяжется с Вами")

	replyKeyboard := tgbotapi.NewRemoveKeyboard(true)
	msg.ReplyMarkup = replyKeyboard

	return &msg
}
