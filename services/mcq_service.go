package services

import (
	"github.com/souvik150/BattleQuiz-Backend/database"
	"github.com/souvik150/BattleQuiz-Backend/models"
)

func CreateMCQ(mcq *models.MCQ) error {
    return database.DB.Create(mcq).Error
}

func GetAllMCQs() []models.MCQ {
    var mcqs []models.MCQ
    database.DB.Find(&mcqs)
    return mcqs
}

func UpdateMCQ(id string, input models.MCQ) error {
    var mcq models.MCQ
    if err := database.DB.Where("id = ?", id).First(&mcq).Error; err != nil {
        return err
    }

    mcq.Question = input.Question
    mcq.Options = input.Options
    mcq.CorrectAnswer = input.CorrectAnswer
    mcq.Difficulty = input.Difficulty

    return database.DB.Save(&mcq).Error
}

func DeleteMCQ(id string) error {
    return database.DB.Where("id = ?", id).Delete(&models.MCQ{}).Error
}
