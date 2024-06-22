package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/souvik150/BattleQuiz-Backend/controllers"
	"github.com/souvik150/BattleQuiz-Backend/middlewares"
)

func RegisterRoutes(r *gin.Engine) {
	auth := r.Group("/auth")
	{
		auth.POST("/signup", controllers.SignUp)
		auth.POST("/login", controllers.Login)
	}

	user := r.Group("/user")
	user.Use(middlewares.AuthMiddleware())
	{
		user.GET("/", controllers.GetUser)
		// user.PUT("/", controllers.UpdateUser)
	}

	games := r.Group("/games")
	games.Use(middlewares.AuthMiddleware())
	{
		games.POST("/", controllers.CreateGame)
		games.GET("/", controllers.GetAllGames)
		games.GET("/:id", controllers.GetGameByID)
		games.PUT("/publish/:id", controllers.PublishGame)
	}

	mcqs := r.Group("/quiz")
	mcqs.Use(middlewares.AuthMiddleware())
	{
		mcqs.POST("/", controllers.CreateQuiz)
		mcqs.GET("/:game_id", controllers.GetQuizzes)
	}

	leaderboard := r.Group("/leaderboard")
	leaderboard.Use(middlewares.AuthMiddleware())
	{
		leaderboard.POST("/", controllers.UpdateLeaderboard)
		leaderboard.GET("/:game_id", controllers.GetLeaderboard)
	}
}
