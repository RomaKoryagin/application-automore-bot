package services

import (
	"alex.com/application-bot/internal/domain/entities"
)

type UserRepository interface {
	GetByChatId(chatId int64) (*entities.User, error)
	Create(chatId int64) (*int, error)
}

type UserService struct {
	UserRepository UserRepository
}

func (userService UserService) GetByChatId(chatId int64) (*entities.User, error) {
	return userService.UserRepository.GetByChatId(chatId)
}

func (userService UserService) CreateByChatId(chatId int64) (*int, error) {
	return userService.UserRepository.Create(chatId)
}

func NewUserService(userRepository UserRepository) *UserService {
	return &UserService{UserRepository: userRepository}
}
