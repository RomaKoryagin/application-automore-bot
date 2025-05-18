package strategies

import (
	"alex.com/application-bot/internal/application/enums"
	"alex.com/application-bot/internal/application/factories"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type CountryReturnCommandStrategy struct {
}

func (strategy *CountryReturnCommandStrategy) Handle(chatId int64, telegramId string, text string) (*tgbotapi.MessageConfig, error) {
	return factories.NewCountryStepMessageFactory().CreateMessage(chatId), nil
}

func (strategy *CountryReturnCommandStrategy) GetKey() enums.StrategyType {
	return enums.CountryReturn
}

func NewCountryReturnCommandStrategy() *CountryReturnCommandStrategy {
	return &CountryReturnCommandStrategy{}
}
