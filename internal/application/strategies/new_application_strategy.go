package strategies

import (
	"log"

	"alex.com/application-bot/internal/application/enums"
	"alex.com/application-bot/internal/application/factories"
	"alex.com/application-bot/internal/domain/entities"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type INewApplicationStrategyUserService interface {
	GetByChatId(chatId int64) (*entities.User, error)
}

type INewApplicationStrategyApplicationService interface {
	GetLastByUserId(userId int) (*entities.Application, error)
	CreateEmptyApplication(userId int) error
}

type NewApplicationStrategy struct {
	UserService        INewApplicationStrategyUserService
	ApplicationService INewApplicationStrategyApplicationService
}

func (strategy NewApplicationStrategy) Handle(chatId int64, text string) (*tgbotapi.MessageConfig, error) {
	user, err := strategy.UserService.GetByChatId(chatId)
	if err != nil {
		log.Printf("error while trying to get user by chat_id, more: %s", err)
		return nil, err
	}

	err = strategy.ApplicationService.CreateEmptyApplication(user.ID)
	if err != nil {
		log.Printf("error while trying to create new empty application, more: %s", err)
		return nil, err
	}

	return factories.CountryStepMessageFactory{}.CreateMessage(chatId), nil
}

func (strategy NewApplicationStrategy) GetKey() enums.StrategyType {
	return enums.NewApplication
}

func NewNewApplicationStrategy(
	userService INewApplicationStrategyUserService,
	applicationService INewApplicationStrategyApplicationService,
) *NewApplicationStrategy {
	return &NewApplicationStrategy{UserService: userService, ApplicationService: applicationService}
}
