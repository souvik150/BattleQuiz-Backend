package models

import (
	"github.com/google/uuid"
)

type Leaderboard struct {
	ID  uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	UserID uuid.UUID `gorm:"type:uuid;not null"`
	GameID uuid.UUID `gorm:"type:uuid;not null"`
	Score  int  `gorm:"not null"`
	Game   Game
	User   User
}