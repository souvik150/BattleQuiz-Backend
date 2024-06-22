package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/souvik150/BattleQuiz-Backend/database"
	"github.com/souvik150/BattleQuiz-Backend/models"
)

type CreateQuizInput struct {
	GameID        uuid.UUID `json:"game_id" binding:"required"`
	Question      string    `json:"question" binding:"required"`
	Options       []string  `json:"options" binding:"required"`
	CorrectAnswer int       `json:"correct_answer" binding:"required"`
	Points        int       `json:"points" binding:"required"`
	Difficulty    string    `json:"difficulty"`
}

type CreateQuizResponse struct {
	ID            uuid.UUID `json:"id"`
	GameID        uuid.UUID `json:"game_id"`
	Question      string    `json:"question"`
	Options       []string  `json:"options"`
	CorrectAnswer int       `json:"correct_answer"`
	Points        int       `json:"points"`
	Difficulty    string    `json:"difficulty"`
}

func CreateQuiz(c *gin.Context) {
	var input CreateQuizInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"error":  err.Error(),
		})
		return
	}

	fmt.Printf("Parsed input: %+v\n", input)

	// Check if the game exists
	var game models.Game
	if err := database.DB.First(&game, "id = ?", input.GameID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Game not found"})
		return
	}

	fmt.Println(game)

	// Create the quiz
	quiz := models.MCQ{
		GameID:        input.GameID,
		Question:      input.Question,
		Options:       input.Options,
		CorrectAnswer: input.CorrectAnswer,
		Points:        input.Points,
		Difficulty:    input.Difficulty,
	}
	if err := database.DB.Create(&quiz).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create quiz"})
		return
	}

	response := CreateQuizResponse{
		ID:            quiz.ID,
		GameID:        quiz.GameID,
		Question:      quiz.Question,
		Options:       quiz.Options,
		CorrectAnswer: quiz.CorrectAnswer,
		Points:        quiz.Points,
		Difficulty:    quiz.Difficulty,
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Quiz created successfully",
		"data":    response,
	})
}

func GetQuizzes(c *gin.Context) {
	gameIdStr := c.Param("game_id")
	gameId, err := uuid.Parse(gameIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid game ID"})
		return
	}

	var quizzes []models.MCQ
	if err := database.DB.Where("game_id = ?", gameId).Find(&quizzes).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Quizzes not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": quizzes})
}
