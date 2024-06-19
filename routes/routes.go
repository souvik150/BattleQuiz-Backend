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

    mcqs := r.Group("/mcqs")
    mcqs.Use(middlewares.AuthMiddleware())
    {
        mcqs.POST("/", controllers.CreateMCQ)
        mcqs.GET("/", controllers.GetMCQs)
        mcqs.PUT("/:id", controllers.UpdateMCQ)
        mcqs.DELETE("/:id", controllers.DeleteMCQ)
    }
}
