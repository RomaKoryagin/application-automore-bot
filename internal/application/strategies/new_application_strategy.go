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
	CreateByChatId(chatId int64) (*int, error)
}

type INewApplicationStrategyApplicationService interface {
	GetLastByUserId(userId int) (*entities.Application, error)
	CreateEmptyApplication(userId int, chatId int64, telegramId string) error
}

type NewApplicationStrategy struct {
	UserService        INewApplicationStrategyUserService
	ApplicationService INewApplicationStrategyApplicationService
}

func (strategy NewApplicationStrategy) Handle(chatId int64, telegramId string, text string) (*tgbotapi.MessageConfig, error) {
	user, err := strategy.UserService.GetByChatId(chatId)
	if err != nil {
		log.Printf("error while trying to get user by chat_id, more: %s", err)
		return nil, err
	}

	if user == nil {
		userId, err := strategy.UserService.CreateByChatId(chatId)
		if err != nil {
			log.Printf("error while trying to create new user, more: %s", err)
			return nil, err
		}

		err = strategy.ApplicationService.CreateEmptyApplication(*userId, chatId, telegramId)
		if err != nil {
			log.Printf("error while trying to create new empty application, more: %s", err)
			return nil, err
		}
	} else {
		err = strategy.ApplicationService.CreateEmptyApplication(user.ID, chatId, telegramId)
		if err != nil {
			log.Printf("error while trying to create new empty application, more: %s", err)
			return nil, err
		}
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
