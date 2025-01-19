package services

import (
	"log"
	"os"
	"time"

	"github.com/Amotekun/payment-system/internal/models"
	"github.com/Amotekun/payment-system/internal/repositories"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func CreateUser(user *models.User) error {
	return repositories.InsertUser(user)
}

func GetUserByEmailOrUsername(email, username string) (*models.User, error) {
	return repositories.FindUsersByEmailOrUsername(email, username)
}

func GenerateJWT(user *models.User) (string, error) {
	claims := jwt.MapClaims{
		"sub":   user.ID,
		"email": user.Email,
		"role":  user.Role,
		"iat":   time.Now().Unix(),
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtSecret)

	if err != nil {
		log.Println("Error signing token:", err)
		return "", err
	}

	return signedToken, nil
}
