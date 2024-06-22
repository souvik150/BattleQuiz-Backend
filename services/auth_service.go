package services

import (
	"github.com/google/uuid"

	"github.com/souvik150/BattleQuiz-Backend/database"
	"github.com/souvik150/BattleQuiz-Backend/models"
)

func CreateUser(user *models.User) error {
    return database.DB.Create(user).Error
}

func GetUserByEmail(email string) (*models.User, error) {
    var user models.User
    err := database.DB.Where("email = ?", email).First(&user).Error
    return &user, err
}

func GetUserById(id uuid.UUID) (*models.User, error) {
    var user models.User
    err := database.DB.Where("id = ?", id).First(&user).Error
    return &user, err
}