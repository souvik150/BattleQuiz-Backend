package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pusher/pusher-http-go/v5"

	"github.com/souvik150/BattleQuiz-Backend/database"
	"github.com/souvik150/BattleQuiz-Backend/models"
	"github.com/souvik150/BattleQuiz-Backend/services"
)

type CreateGameInput struct {
	Title string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Image string `json:"image" binding:"required"`
}

type CreateGameResponse struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
}

type GetGameResponse struct {
	ID          uuid.UUID   `json:"id"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Image       string      `json:"image"`
	CreatedBy   GetUserResponse `json:"created_by"`
	Quizzes		 []GetQuizResponse `json:"quizzes"`
}

type GetUserResponse struct {
	ID       uuid.UUID `json:"id"`
	FullName string    `json:"fullname"`
	Username string    `json:"username"`
}

type GetQuizResponse struct {
	ID         uuid.UUID `json:"id"`
	Question   string    `json:"question"`
	Options    []string  `json:"options"`
	Points     int       `json:"points"`
	Difficulty string    `json:"difficulty"`
}

type LeaderboardInput struct {
	GameID uint `json:"game_id" binding:"required"`
	UserID uint `json:"user_id" binding:"required"`
	Score  int  `json:"score" binding:"required"`
}

var pusherClient = pusher.Client{
	AppID:   "1823253",
	Key:     "849f5b01c35dfd0fec56",
	Secret:  "52eae2220b2814b4866e",
	Cluster: "ap2",
	Secure:  true,
}

func CreateGame(c *gin.Context) {
	var input CreateGameInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	fmt.Println(userID)

	// Convert userID to uuid.UUID
	userUUID, err := uuid.Parse(userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID"})
		return
	}

	game := models.Game{
		ID: 				uuid.New(),
		Title:       input.Title,
		Description: input.Description,
		Image:       input.Image,
		UserID:      userUUID,
	}
	if err := services.CreateGame(&game); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create game"})
		return
	}
	
	var gameResponse CreateGameResponse
	gameResponse.ID = game.ID
	gameResponse.Title = game.Title
	gameResponse.Description = game.Description
	gameResponse.Image = game.Image

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Game created successfully",
		"data":    gameResponse,
	})
}

func GetAllGames(c *gin.Context) {
	var games []models.Game
	
	if err := database.DB.Where("is_live = ?", true).Preload("CreatedBy").Find(&games).Error; 
	
	err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Games not found"})
		return
	}

	if len(games) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"status": true,
			"message": "No games found",
			"data": nil,
		})
		return
	}

	var gameResponses []GetGameResponse
	for _, game := range games {
		response := GetGameResponse{
			ID:          game.ID,
			Title:       game.Title,
			Description: game.Description,
			Image:       game.Image,
			CreatedBy: GetUserResponse{
				ID:       game.CreatedBy.ID,
				FullName: game.CreatedBy.FullName,
				Username: game.CreatedBy.Username,
			},
		}
		gameResponses = append(gameResponses, response)
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"message": "Games fetched successfully",
		"data": gameResponses,
	})
}



func GetGameByID(c *gin.Context) {
	var game models.Game

	gameIdStr := c.Param("id")
	gameId, err := uuid.Parse(gameIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid game ID"})
		return
	}

	if err := database.DB.Preload("CreatedBy").Preload("Quizzes").First(&game, "id = ?", gameId).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Game not found"})
		return
	}

	var quizResponses []GetQuizResponse
	for _, quiz := range game.Quizzes {
		quizResponses = append(quizResponses, GetQuizResponse{
			ID:         quiz.ID,
			Question:   quiz.Question,
			Options:    quiz.Options,
			Points:     quiz.Points,
			Difficulty: quiz.Difficulty,
		})
	}

	gameResponse := GetGameResponse{
		ID:          game.ID,
		Title:       game.Title,
		Description: game.Description,
		Image:       game.Image,
		CreatedBy: GetUserResponse{
			ID:       game.CreatedBy.ID,
			FullName: game.CreatedBy.FullName,
			Username: game.CreatedBy.Username,
		},
		Quizzes: quizResponses,
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Game fetched successfully",
		"data":    gameResponse,
	})
}
func PublishGame(c *gin.Context) {
	var game models.Game

	gameIdStr := c.Param("id")
	gameId, err := uuid.Parse(gameIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid game ID"})
		return
	}

	if err := database.DB.First(&game, "id = ?", gameId).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Game not found"})
		return
	}

	if game.IsLive {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "Game already published",
		})
		return
	}

	game.IsLive = true
	if err := database.DB.Save(&game).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to publish game"})
		return
	}

	pusherClient.Trigger("game-channel", "game-published", game)
	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Game published successfully",
	})
}

func UpdateLeaderboard(c *gin.Context) {
	// var input LeaderboardInput
	// if err := c.ShouldBindJSON(&input); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	// var leaderboard models.Leaderboard
	// if err := database.DB.Where("game_id = ? AND user_id = ?", input.GameID, input.UserID).First(&leaderboard).Error; err != nil {
	// 	leaderboard = models.Leaderboard{
	// 		GameID: input.GameID,
	// 		UserID: input.UserID,
	// 		Score:  input.Score,
	// 	}
	// 	if err := database.DB.Create(&leaderboard).Error; err != nil {
	// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update leaderboard"})
	// 		return
	// 	}
	// } else {
	// 	leaderboard.Score = input.Score
	// 	if err := database.DB.Save(&leaderboard).Error; err != nil {
	// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update leaderboard"})
	// 		return
	// 	}
	// }

	c.JSON(http.StatusOK, gin.H{"data": "leaderboard"})
}

func GetLeaderboard(c *gin.Context) {
	var leaderboard []models.Leaderboard
	if err := database.DB.Where("game_id = ?", c.Param("game_id")).Order("score desc").Preload("User").Find(&leaderboard).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Leaderboard not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": leaderboard})
}
