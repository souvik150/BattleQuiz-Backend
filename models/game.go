package models

import (
	"github.com/google/uuid"
)

type Game struct {
	ID          uuid.UUID     `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Title       string        `gorm:"not null"`
	Description string        `gorm:"not null"`
	Image       string        `gorm:"default:'https://www.elevenforum.com/data/attachments/12/12237-4d23287149b3f4ebd9777653225df31b.jpg'"`
	IsLive      bool          `gorm:"default:false"`
	Quizzes     []MCQ         `gorm:"foreignkey:GameID"`
	Leaderboards []Leaderboard `gorm:"foreignkey:GameID"`
  UserID       uuid.UUID     `gorm:"type:uuid;not null"`
  CreatedBy    User          `gorm:"foreignkey:UserID"`
}
