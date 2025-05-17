package strategies

import (
	"database/sql"
	"log"
	"regexp"

	"alex.com/application-bot/internal/application/enums"
	"alex.com/application-bot/internal/application/factories"
	"alex.com/application-bot/internal/domain/entities"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type IUpdateApplicationStrategyUserService interface {
	GetByChatId(chatId int64) (*entities.User, error)
}

type IUpdateApplicationApplicationService interface {
	GetLastByUserId(userId int) (*entities.Application, error)
	GetStepTypeByNumber(appl *entities.Application) enums.StepType
	Update(appl *entities.Application) error
}

type IUpdateApplicationFactoryResolver interface {
	Resolve(fType enums.StepType) (factories.MessageFactory, error)
}

type UpdateApplicationStategy struct {
	UserService        IUpdateApplicationStrategyUserService
	ApplicationService IUpdateApplicationApplicationService
	FactoryResolver    IUpdateApplicationFactoryResolver
}

func (strategy UpdateApplicationStategy) Handle(chatId int64, telegramId string, text string) (*tgbotapi.MessageConfig, error) {
	user, err := strategy.UserService.GetByChatId(chatId)
	if err != nil {
		log.Printf("error while getting user by chat_id, more: %s", err)
		return nil, err
	}

	appl, err := strategy.ApplicationService.GetLastByUserId(user.ID)
	if err != nil {
		log.Printf("error while getting last application by user_id, more: %s", err)
		return nil, err
	}

	if appl == nil {
		// @TODO create appliction or think what should we do in this case
	}

	stepType := strategy.ApplicationService.GetStepTypeByNumber(appl)
	var value sql.NullString
	value.Valid = true
	value.String = text
	switch stepType {
	case enums.Country:
		appl.Country = value
		if value.String == "Япония" {
			stepType = enums.JapanWarning
		}
	case enums.City:
		appl.City = value
	case enums.Budget:
		appl.Budget = value
	case enums.JapanWarning:
		if text != "Вернуться к выбору страны" {
			appl.SteeringWheelType = value
			stepType = enums.City
		} else {
			value.Valid = false
			appl.SteeringWheelType = value
			appl.Country = value
		}
	case enums.MarkOrConditions:
		appl.MarkOrConditions = value
	case enums.PersonPhone:
		appl.PersonPhone = value
	case enums.PersonName:
		appl.PersonName = value
	}

	isValid := true
	if stepType == enums.PersonPhone {
		pattern := regexp.MustCompile(`^7\d{10}$`)
		isValid = pattern.MatchString(text)
	}

	var resultStep enums.StepType
	if isValid {
		if stepType != enums.JapanWarning {
			appl.Step++
		}
		err = strategy.ApplicationService.Update(appl)
		if err != nil {
			log.Printf("error while trying to update application, more: %s", err)
			return nil, err
		}
		resultStep = strategy.ApplicationService.GetStepTypeByNumber(appl)
	} else {
		resultStep = enums.InvalidPhone
	}

	factory, err := strategy.FactoryResolver.Resolve(resultStep)
	if err != nil {
		return nil, err
	}

	msg := factory.CreateMessage(chatId)

	return msg, nil
}

func (strategy UpdateApplicationStategy) GetKey() enums.StrategyType {
	return enums.UpdateApplication
}

func NewUpdateApplicationStrategy(
	userService IUpdateApplicationStrategyUserService,
	applicationService IUpdateApplicationApplicationService,
	factoryResolver IUpdateApplicationFactoryResolver,
) *UpdateApplicationStategy {
	return &UpdateApplicationStategy{
		UserService:        userService,
		ApplicationService: applicationService,
		FactoryResolver:    factoryResolver,
	}
}
