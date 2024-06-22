package services

import (
	"github.com/souvik150/BattleQuiz-Backend/database"
	"github.com/souvik150/BattleQuiz-Backend/models"
)

func CreateGame(game *models.Game) error {
	if err := database.DB.Create(game).Error; err != nil {
			return err
	}
	return nil
}