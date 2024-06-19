package database

import (
	"github.com/souvik150/BattleQuiz-Backend/models"
)

func Migrate() {
    DB.AutoMigrate(&models.User{}, &models.MCQ{})
}
