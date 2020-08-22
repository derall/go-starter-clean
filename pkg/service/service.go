package service

import (
	"go-starter-clean/pkg/entity"
)

type userService struct {
	userRepo entity.UserRepository
}

func New(ur entity.UserRepository) entity.UserService {
	return &userService{
		userRepo: ur,
	}
}
func (service *userService) Store(user *entity.User) error {
	return service.userRepo.Store(user)
}

func (service *userService) FindAll() ([]*entity.User, error) {
	return nil, nil
}
