package models

import (
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

type MCQ struct {
	gorm.Model
	Question        string         `gorm:"not null"`
	Options         pq.StringArray `gorm:"type:text[]"`
	CorrectAnswer   string         `gorm:"not null"`
	DifficultyLevel string
}
