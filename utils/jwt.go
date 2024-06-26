package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"

	"github.com/souvik150/BattleQuiz-Backend/config"
)

func GenerateToken(userID uuid.UUID) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": userID,
        "exp":     time.Now().Add(time.Hour * 72).Unix(),
    })

    return token.SignedString([]byte(config.GetJWTSecret()))
}
