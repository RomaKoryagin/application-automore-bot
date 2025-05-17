package strategies

import (
	"log"

	"alex.com/application-bot/internal/application/enums"
	"alex.com/application-bot/internal/domain/entities"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type IMessageStrategy interface {
	Handle(chatId int64, telegramId string, text string) (*tgbotapi.MessageConfig, error)

	GetKey() enums.StrategyType
}

type IStrategyResolverApplicationService interface {
	GetLastByChatId(chatId int64) (*entities.Application, error)
}

type StrategyResolver struct {
	strategies         map[enums.StrategyType]IMessageStrategy
	ApplicationService IStrategyResolverApplicationService
}

func (resolver StrategyResolver) AddStrategy(strategy IMessageStrategy) {
	resolver.strategies[strategy.GetKey()] = strategy
}

func (resolver StrategyResolver) Resolve(chatId int64, text string) IMessageStrategy {
	var sType enums.StrategyType = enums.UpdateApplication
	appl, err := resolver.ApplicationService.GetLastByChatId(chatId)

	if err != nil {
		sType = enums.Error
	}

	if appl != nil && appl.Step >= 7 {
		sType = enums.NoActiveApplication
	}

	if text == "/menu" {
		sType = enums.ShowMenu
	}

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

	log.Println("type " + sType)

	return resolver.strategies[sType]
}

func NewStrategyResolver(applicationService IStrategyResolverApplicationService) *StrategyResolver {
	return &StrategyResolver{
		strategies:         make(map[enums.StrategyType]IMessageStrategy),
		ApplicationService: applicationService,
	}
}
