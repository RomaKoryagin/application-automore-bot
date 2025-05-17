package strategies

import (
	"alex.com/application-bot/internal/application/enums"
	"alex.com/application-bot/internal/application/factories"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type NoActiveApplicationStrategy struct {
}

func (strategy NoActiveApplicationStrategy) Handle(chatId int64, telegramId string, text string) (*tgbotapi.MessageConfig, error) {
	return factories.NewNoActiveApplicationMessageFactory().CreateMessage(chatId), nil
}

func (strategy NoActiveApplicationStrategy) GetKey() enums.StrategyType {
	return enums.NoActiveApplication
}

func NewNoActiveApplicationStrategy() *NoActiveApplicationStrategy {
	return &NoActiveApplicationStrategy{}
}
