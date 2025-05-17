package builders

import (
	"fmt"

	"alex.com/application-bot/internal/domain/entities"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type ApplicationMessageBuilder struct {
	Msg *tgbotapi.MessageConfig
}

func (builder *ApplicationMessageBuilder) ConfigureChannelName(channelName string) *ApplicationMessageBuilder {
	builder.Msg.ChannelUsername = channelName
	return builder
}

func (builder *ApplicationMessageBuilder) ConfigureApplicationText(appl *entities.Application) *ApplicationMessageBuilder {
	text := "<b>Новая заявка</b>"

	text += fmt.Sprintf("\nИмя: %s ", appl.PersonName.String)

	text += fmt.Sprintf("\nНомер телефона: +%s", appl.PersonPhone.String)

	text += fmt.Sprintf("\n\nТелегам ID: @%s", appl.TelegramId)

	text += fmt.Sprintf("\nГород: %s", appl.City.String)

	text += fmt.Sprintf("\nАвтомобиль: %s", appl.MarkOrConditions.String)

	text += fmt.Sprintf("\nСтрана: %s", appl.Country.String)

	text += fmt.Sprintf("\nБюджет: %s", appl.Budget.String)

	text += fmt.Sprintf("\nID: %d", appl.ChatId)

	builder.Msg.Text = text

	return builder
}

func (builder *ApplicationMessageBuilder) AddOpenChatButton(username string) *ApplicationMessageBuilder {
	button := tgbotapi.NewInlineKeyboardButtonURL("Перейти к диалогу в TG", fmt.Sprintf("https://t.me/%s", username))

	replyKeyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(button),
	)
	builder.Msg.ReplyMarkup = replyKeyboard
	return builder
}

func (builder *ApplicationMessageBuilder) Build() *tgbotapi.MessageConfig {
	builder.Msg.ParseMode = "HTML"
	return builder.Msg
}

func NewApplicationMessageBuilder() *ApplicationMessageBuilder {
	msg := tgbotapi.MessageConfig{}
	return &ApplicationMessageBuilder{Msg: &msg}
}
