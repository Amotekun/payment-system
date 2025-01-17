package services

import (
	"github.com/Amotekun/payment-system/internal/models"
	"github.com/Amotekun/payment-system/internal/repositories"
)

func CreateUser(user *models.User) error {
	return repositories.InsertUser(user)
}

func GetUserByEmailOrUsername(email, username string) (*models.User, error) {
	return repositories.FindUsersByEmailOrUsername(email, username)
}
