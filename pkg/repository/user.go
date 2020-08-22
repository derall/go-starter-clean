package repository

import (
	"errors"
	"fmt"
	"go-starter-clean/pkg/entity"
	"go-starter-clean/pkg/logger"

	"github.com/jinzhu/gorm"
)

type userRepository struct {
	DB  *gorm.DB
	log logger.LogInfoFormat
}

func NewUserRepository(db *gorm.DB, log logger.LogInfoFormat) entity.UserRepository {
	return &userRepository{
		DB:  db,
		log: log,
	}
}

func (mu *userRepository) GetAll() ([]*entity.User, error) {
	return nil, nil
}

func (mu *userRepository) Store(user *entity.User) error {
	mu.log.Debugf("creating the user with email : %v", user.Email)
	err := mu.DB.Create(&user)
	if err != nil {
		mu.log.Errorf("error while creating the user, reason : %v", err)
		return err.Error
	}
	return nil
}

func (mu *userRepository) Delete(id string) error {
	mu.log.Debugf("deleting the user with id : %s", id)

	if mu.DB.Delete(&entity.User{}, "id = ?", id).Error != nil {
		errMsg := fmt.Sprintf("error while deleting the user with id : %s", id)
		mu.log.Errorf(errMsg)
		return errors.New(errMsg)
	}
	return nil

}

func (mu *userRepository) GetByID(id string) (*entity.User, error) {
	mu.log.Debugf("get user details by id : %s", id)

	user := &entity.User{}
	err := mu.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		mu.log.Errorf("user not found with id : %s, reason : %v", id, err)
		return nil, err
	}
	return user, nil
}
