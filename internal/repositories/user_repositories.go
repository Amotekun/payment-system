package repositories

import (
	"github.com/Amotekun/payment-system/internal/config"
	"github.com/Amotekun/payment-system/internal/models"
)

func InsertUser(user *models.User) error {
	return config.DB.Create(user).Error
}

func FindUsersByEmailOrUsername(email, username string) (*models.User, error) {
	var user models.User
	result := config.DB.Where("email = ? OR username = ?", email, username).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
