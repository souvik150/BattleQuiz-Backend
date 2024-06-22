package database

import (
	"github.com/souvik150/BattleQuiz-Backend/models"
)

func Migrate() {
	err := DB.AutoMigrate(&models.User{}, &models.MCQ{}, &models.Game{}, &models.Leaderboard{}).Error
	if err != nil {
		print("Failed to migrate database", err)
		panic(err)
	}
}
