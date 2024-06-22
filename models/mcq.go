package models

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type MCQ struct {
	ID            uuid.UUID      `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	GameID        uuid.UUID      `gorm:"type:uuid;not null"`
	Question      string         `gorm:"not null"`
	Options       pq.StringArray `gorm:"type:text[]"`
	CorrectAnswer int            `gorm:"type:int"`
	Points        int            `gorm:"not null"`
	Difficulty    string
}
