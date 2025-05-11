package services

import (
	"alex.com/application-bot/internal/application/enums"
	"alex.com/application-bot/internal/application/strategies"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type IMessageStrategy interface {
	Handle(chatId int64, text string) (*tgbotapi.MessageConfig, error)

	GetKey() enums.StrategyType
}

type IStrategyResolver interface {
	Resolve(chatId int64, text string) strategies.IMessageStrategy
}

type TelegramMessageService struct {
	Bot              *tgbotapi.BotAPI
	StrategyResolver IStrategyResolver
}

func (service TelegramMessageService) SendReplyMessage(chatId int64, text string) {
	strategy := service.StrategyResolver.Resolve(chatId, text)

	msg, _ := strategy.Handle(chatId, text)

	service.Bot.Send(msg)
}

func NewTelegramMessageService(
	bot *tgbotapi.BotAPI,
	strategyResolver IStrategyResolver,
) *TelegramMessageService {

	service := TelegramMessageService{Bot: bot, StrategyResolver: strategyResolver}
	return &service
}
