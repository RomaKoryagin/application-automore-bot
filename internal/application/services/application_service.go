package services

import (
	"log"

	"alex.com/application-bot/internal/application/enums"
	"alex.com/application-bot/internal/domain/entities"
	"alex.com/application-bot/internal/infrastructure/repositories"
)

type ApplicationService struct {
	ApplicationRepository *repositories.ApplicationRepository
	UserRepository        *repositories.UserRepository
}

func (service ApplicationService) GetLastByUserId(userId int) (*entities.Application, error) {
	return service.ApplicationRepository.GetLastByUserId(userId)
}

func (service ApplicationService) GetLastByChatId(chatId int64) (*entities.Application, error) {
	user, err := service.UserRepository.GetByChatId(chatId)
	if err != nil {
		log.Printf("error while getting user by chat_id, more: %s", err)
		return nil, err
	}

	if user == nil {
		return nil, nil
	}

	appl, err := service.ApplicationRepository.GetLastByUserId(user.ID)
	if err != nil {
		log.Printf("error while getting last application by user_id, more: %s", err)
		return nil, err
	}

	if appl == nil {
		log.Printf("there is no last application for user with id: %d", user.ID)
		return nil, nil
	}

	return appl, nil
}

func (service ApplicationService) GetStepTypeByNumber(appl *entities.Application) enums.StepType {
	stepNumberToTypeMap := map[int]enums.StepType{
		1: enums.Country,
		2: enums.City,
		3: enums.Budget,
		4: enums.MarkOrConditions,
		5: enums.PersonName,
		6: enums.PersonPhone,
		7: enums.Submit,
		8: enums.SubmittedApplication,
	}

	step := stepNumberToTypeMap[appl.Step]

	if step == enums.Country && appl.Country.Valid && appl.Country.String == "Япония" {
		step = enums.JapanWarning
	}

	if step == enums.Country && appl.Country.Valid && appl.Country.String != "Япония" {
		step = enums.City
	}

	return step
}

func (service ApplicationService) CreateEmptyApplication(userId int) error {
	return service.ApplicationRepository.CreateEmpty(userId)
}

func (service ApplicationService) Update(appl *entities.Application) error {
	return service.ApplicationRepository.Update(appl)
}

func NewApplicationService(
	applicationRepository *repositories.ApplicationRepository,
	userRepository *repositories.UserRepository,
) *ApplicationService {
	return &ApplicationService{
		ApplicationRepository: applicationRepository,
		UserRepository:        userRepository,
	}
}
