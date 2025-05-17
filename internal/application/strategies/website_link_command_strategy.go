package strategies

import (
	"alex.com/application-bot/internal/application/enums"
	"alex.com/application-bot/internal/application/factories"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type WebsiteCommandStategy struct {
}

func (strategy WebsiteCommandStategy) Handle(chatId int64, telegramId string, text string) (*tgbotapi.MessageConfig, error) {
	return factories.WebsiteLinkMessageFactory{}.CreateMessage(chatId), nil
}

func (strategy WebsiteCommandStategy) GetKey() enums.StrategyType {
	return enums.WebsiteLink
}

func NewWebsiteCommandStrategy() *WebsiteCommandStategy {
	return &WebsiteCommandStategy{}
}
