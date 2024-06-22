package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID     `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
    FullName string `gorm:"not null"`
    Username string `gorm:"unique;not null"`
    Email    string `gorm:"unique;not null"`
    Password string `gorm:"not null"`
    Image    string `gorm:"default:'https://w7.pngwing.com/pngs/981/645/png-transparent-default-profile-united-states-computer-icons-desktop-free-high-quality-person-icon-miscellaneous-silhouette-symbol-thumbnail.png'"`
    Active   bool   `gorm:"default:true"`
    CreatedAt time.Time
    UpdatedAt time.Time
}