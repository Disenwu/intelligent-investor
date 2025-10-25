package store

import (
	"intelligent-investor/internal/app/model"
	config "intelligent-investor/internal/pkg/service"
)

type UserStore interface {
	CreateUser(user *model.User) error
	GetUserByUsername(username string) (*model.User, error)
}

func CreateUser(user *model.User) error {
	return config.DB.Create(user).Error
}

func GetUserByUsername(username string) (*model.User, error) {
	var user model.User
	if err := config.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	if err := config.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
