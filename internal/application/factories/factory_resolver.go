package factories

import (
	"errors"

	"alex.com/application-bot/internal/application/enums"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type MessageFactory interface {
	CreateMessage(chatId int64) *tgbotapi.MessageConfig
}

type FactoryResolver struct {
}

func (resolver FactoryResolver) Resolve(fType enums.StepType) (MessageFactory, error) {
	switch fType {
	case enums.Budget:
		return &BudgetStepMessageFactory{}, nil
	case enums.City:
		return &CityStepMessageFactory{}, nil
	case enums.Country:
		return &CountryStepMessageFactory{}, nil
	case enums.JapanWarning:
		return &JapanWarningMessageFactory{}, nil
	case enums.MarkOrConditions:
		return &MarkOfConditionsStepMessageFactory{}, nil
	case enums.PersonName:
		return &PersonNameStepMessageFactory{}, nil
	case enums.PersonPhone:
		return &PersonPhoneStepMessageFactory{}, nil
	case enums.InvalidPhone:
		return &InvalidPhoneMessageFactory{}, nil
	case enums.SubmittedApplication:
		return &SubmittedApplicationMessageFactory{}, nil
	case enums.Menu:
		return &MenuMessageFactory{}, nil
	default:
		return nil, errors.New("incorrent factory type found")
	}
}

func NewFactoryResolver() *FactoryResolver {
	return &FactoryResolver{}
}
