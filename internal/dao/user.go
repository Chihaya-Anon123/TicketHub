package dao

import (
	"errors"

	"github.com/Chihaya-Anon123/TicketHub/internal/database"
	"github.com/Chihaya-Anon123/TicketHub/internal/model"
	"gorm.io/gorm"
)

// 创建用户
func CreateUser(user *model.User) error {
	return database.DB.Create(user).Error
}

// 通过 Username 查找用户
func GetUserByUsername(username string) (*model.User, error) {
	var user model.User
	err := database.DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// 通过 Email 查找用户
func GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	err := database.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
