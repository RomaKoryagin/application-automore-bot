package strategies

import (
	"alex.com/application-bot/internal/application/enums"
	"alex.com/application-bot/internal/application/factories"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type JapanWarningCommandStrategy struct {
}

func (strategy *JapanWarningCommandStrategy) Handle(chatId int64, telegramId string, text string) (*tgbotapi.MessageConfig, error) {
	return factories.NewJapanWarningMessageFactory().CreateMessage(chatId), nil
}

func (strategy *JapanWarningCommandStrategy) GetKey() enums.StrategyType {
	return enums.JapanWheelWarning
}

func NewJapanWarningCommandStrategy() *JapanWarningCommandStrategy {
	return &JapanWarningCommandStrategy{}
}
