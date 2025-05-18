package strategies

import (
	"errors"
	"log"

	"alex.com/application-bot/internal/application/enums"
	"alex.com/application-bot/internal/application/factories"
	"alex.com/application-bot/internal/domain/entities"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type IUpdateCountryStrategyUserService interface {
	GetByChatId(chatId int64) (*entities.User, error)
}

type IUpdateCountryApplicationService interface {
	GetLastByUserId(userId int) (*entities.Application, error)
	GetStepTypeByNumber(appl *entities.Application) enums.StepType
	Update(appl *entities.Application) error
}

type UpdateCountryCommandStrategy struct {
	UserService        IUpdateCountryStrategyUserService
	ApplicationService IUpdateCountryApplicationService
}

func (strategy *UpdateCountryCommandStrategy) Handle(chatId int64, telegramId string, text string) (*tgbotapi.MessageConfig, error) {
	user, err := strategy.UserService.GetByChatId(chatId)

	if err != nil {
		log.Printf("error while getting user by chat_id, more: %s", err)
		return nil, err
	}

	if user == nil {
		return nil, errors.New("error user not found")
	}

	appl, err := strategy.ApplicationService.GetLastByUserId(user.ID)
	if err != nil {
		log.Printf("error while getting last application by user_id, more: %s", err)
		return nil, err
	}

	if appl == nil {
		return nil, errors.New("application now found")
	}

	stepType := strategy.ApplicationService.GetStepTypeByNumber(appl)

	if stepType == enums.Country {

		if text == "/korea" {
			appl.Country.String = "Корея"
			appl.Country.Valid = true

			appl.Step++
		}

		if text == "/china" {
			appl.Country.String = "Китай"
			appl.Country.Valid = true

			appl.Step++
		}

		if text == "/submit-right-wheeling-type" {
			appl.Country.String = "Япония"
			appl.Country.Valid = true

			appl.Step++
		}

		err := strategy.ApplicationService.Update(appl)

		if err != nil {
			log.Printf("error while trying to update application")
		}

		return factories.NewCityStepMessageFactory().CreateMessage(chatId), nil
	}

	return nil, nil
}

func (strategy *UpdateCountryCommandStrategy) GetKey() enums.StrategyType {
	return enums.CountryResolving
}

func NewUpdateCountryStrategy(
	userService IUpdateCountryStrategyUserService,
	applicationService IUpdateCountryApplicationService,
) *UpdateCountryCommandStrategy {
	return &UpdateCountryCommandStrategy{
		UserService:        userService,
		ApplicationService: applicationService,
	}
}
