package strategies

import (
	"alex.com/application-bot/internal/application/enums"
	"alex.com/application-bot/internal/application/factories"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type AboutCommandStrategy struct {
}

func (strategy AboutCommandStrategy) Handle(chatId int64, text string) (*tgbotapi.MessageConfig, error) {
	return factories.AboutMessageFactory{}.CreateMessage(chatId), nil
}

func (strategy AboutCommandStrategy) GetKey() enums.StrategyType {
	return enums.About
}

func NewAboutCommandStrategy() *AboutCommandStrategy {
	return &AboutCommandStrategy{}
}
