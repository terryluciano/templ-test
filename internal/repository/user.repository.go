package repository

import (
	"github.com/terryluciano/templ-test/internal/database"
	"github.com/terryluciano/templ-test/internal/model"
)

func CreateUser(user *model.User) error {
	if err := database.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func GetUserByID(id uint) (*model.User, error) {
	var user model.User
	if err := database.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	if err := database.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
