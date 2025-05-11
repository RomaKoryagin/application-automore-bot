package strategies

import (
	"alex.com/application-bot/internal/application/enums"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type IMessageStrategy interface {
	Handle(chatId int64, text string) (*tgbotapi.MessageConfig, error)

	GetKey() enums.StrategyType
}

type StrategyResolver struct {
	strategies map[enums.StrategyType]IMessageStrategy
}

func (resolver StrategyResolver) AddStrategy(strategy IMessageStrategy) {
	resolver.strategies[strategy.GetKey()] = strategy
}

func (resolver StrategyResolver) Resolve(chatId int64, text string) IMessageStrategy {
	var sType enums.StrategyType = enums.UpdateApplication

	if text == "/start" {
		sType = enums.Start
	}

	if text == "/link" {
		sType = enums.WebsiteLink
	}

	if text == "/newapplication" {
		sType = enums.NewApplication
	}

	if text == "/about" {
		sType = enums.About
	}

	return resolver.strategies[sType]
}

func NewStrategyResolver() *StrategyResolver {
	return &StrategyResolver{
		strategies: make(map[enums.StrategyType]IMessageStrategy),
	}
}
