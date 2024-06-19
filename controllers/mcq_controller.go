package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/souvik150/BattleQuiz-Backend/models"
	"github.com/souvik150/BattleQuiz-Backend/services"
)

type CreateMCQInput struct {
    Question       string   `json:"question" binding:"required"`
    Options        []string `json:"options" binding:"required"`
    CorrectAnswer  string   `json:"correct_answer" binding:"required"`
    DifficultyLevel string   `json:"difficulty_level"`
}

func CreateMCQ(c *gin.Context) {
    var input CreateMCQInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    mcq := models.MCQ{
        Question:       input.Question,
        Options:        input.Options,
        CorrectAnswer:  input.CorrectAnswer,
        DifficultyLevel: input.DifficultyLevel,
    }
    services.CreateMCQ(&mcq)

    c.JSON(http.StatusOK, gin.H{"message": "MCQ created successfully"})
}

func GetMCQs(c *gin.Context) {
    mcqs := services.GetAllMCQs()
    c.JSON(http.StatusOK, gin.H{"mcqs": mcqs})
}

func UpdateMCQ(c *gin.Context) {
    var input CreateMCQInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    id := c.Param("id")
    err := services.UpdateMCQ(id, models.MCQ{
					Question:       input.Question,
					Options:        input.Options,
					CorrectAnswer:  input.CorrectAnswer,
					DifficultyLevel: input.DifficultyLevel,
			})
			
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "MCQ updated successfully"})
}

func DeleteMCQ(c *gin.Context) {
    id := c.Param("id")
    err := services.DeleteMCQ(id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "MCQ deleted successfully"})
}
