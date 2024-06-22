package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/souvik150/BattleQuiz-Backend/models"
	"github.com/souvik150/BattleQuiz-Backend/services"
	"github.com/souvik150/BattleQuiz-Backend/utils"
)

type SignUpInput struct {
    Fullname string `json:"fullname" binding:"required"`
    Username string `json:"username" binding:"required"`
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required"`
}

type UserResponse struct {
    ID       uuid.UUID `json:"id"`
    Fullname string `json:"fullname"`
    Username string `json:"username"`
    Email    string `json:"email"`
    Image    string `json:"image"`
}

type LoginInput struct {
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required"`
}

func SignUp(c *gin.Context) {
    var input SignUpInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    hashedPassword, err := utils.HashPassword(input.Password)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "status": false,
            "message": "Failed to create user",
            "data": nil,
        })
        return
    }

    user := models.User{Username: input.Username, Email: input.Email, Password: hashedPassword, FullName: input.Fullname}
    err = services.CreateUser(&user)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "status": false,
            "message": "Failed to create user",
            "data": err.Error(),
        })
        return
    }

    var userResponse UserResponse
    userResponse.ID = user.ID
    userResponse.Fullname = user.FullName
    userResponse.Username = user.Username
    userResponse.Email = user.Email
    userResponse.Image = user.Image

    c.JSON(http.StatusOK, gin.H{
        "status": true,
        "message": "User created successfully",
        "data": userResponse,
    })
}

func Login(c *gin.Context) {
    var input LoginInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    user, err := services.GetUserByEmail(input.Email)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    if err := utils.CheckPassword(input.Password, user.Password); err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    token, err := utils.GenerateToken(user.ID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "status": true,
        "message": "Logged in successfully",
        "data": gin.H{"token": token},
    })
}

func GetUser(c *gin.Context) {
    userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found"})
		return
	}
    userId, err := uuid.Parse(userID.(string))
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user ID"})
        return
    }

    user, err := services.GetUserById(userId)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

    var userResponse UserResponse
    userResponse.ID = user.ID
    userResponse.Fullname = user.FullName
    userResponse.Username = user.Username
    userResponse.Email = user.Email
    userResponse.Image = user.Image

    c.JSON(http.StatusOK, gin.H{
        "status": true,
        "message": "User fetched successfully",
        "data": userResponse,
    })
}