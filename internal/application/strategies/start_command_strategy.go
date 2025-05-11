package strategies

import (
	"errors"
	"log"

	"alex.com/application-bot/internal/application/enums"
	"alex.com/application-bot/internal/application/factories"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type IApplicationService interface {
	CreateEmptyApplication(userId int) error
}

type IUserService interface {
	CreateByChatId(chatId int64) (*int, error)
}

type StartCommandStategy struct {
	UserService        IUserService
	ApplicationService IApplicationService
}

// @TODO add transaction
func (strategy StartCommandStategy) Handle(chatId int64, text string) (*tgbotapi.MessageConfig, error) {
	userId, err := strategy.UserService.CreateByChatId(chatId)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if userId == nil {
		return nil, errors.New("created user id is undefined")
	}

	err = strategy.ApplicationService.CreateEmptyApplication(*userId)

	if err != nil {
		return nil, err
	}

	return factories.CountryStepMessageFactory{}.CreateMessage(chatId), nil
}

func (strategy StartCommandStategy) GetKey() enums.StrategyType {
	return enums.Start
}

func NewStartCommandStrategy(userService IUserService, applService IApplicationService) StartCommandStategy {
	return StartCommandStategy{UserService: userService, ApplicationService: applService}
}
