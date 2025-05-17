package strategies

import (
	"alex.com/application-bot/internal/application/enums"
	"alex.com/application-bot/internal/application/factories"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type MenuCommandStrategy struct {
}

func (strategy MenuCommandStrategy) Handle(chatId int64, telegramId string, text string) (*tgbotapi.MessageConfig, error) {
	return factories.MenuMessageFactory{}.CreateMessage(chatId), nil
}

func (strategy MenuCommandStrategy) GetKey() enums.StrategyType {
	return enums.ShowMenu
}

func NewMenuCommandStrategy() *MenuCommandStrategy {
	return &MenuCommandStrategy{}
}
