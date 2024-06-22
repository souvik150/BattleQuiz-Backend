package database

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/souvik150/BattleQuiz-Backend/config"
	"github.com/souvik150/BattleQuiz-Backend/models"
)

var DB *gorm.DB

func Connect() error {
	var err error
	DB, err = gorm.Open("postgres", config.GetDBConfig())
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
		return err
	}

	err = DB.AutoMigrate(&models.User{}).Error
	if err != nil {
		print("Failed to migrate user", err)
		panic(err)
	}
	log.Println("User migrated successfully")

	err = DB.AutoMigrate(&models.MCQ{}).Error
	if err != nil {
		print("Failed to migrate mcq", err)
		panic(err)
	}
	log.Println("MCQ migrated successfully")

	err = DB.AutoMigrate(&models.Game{}).Error
	if err != nil {
		print("Failed to migrate game", err)
		panic(err)
	}
	log.Println("Game migrated successfully")

	err = DB.AutoMigrate(&models.Leaderboard{}).Error
	if err != nil {
		print("Failed to migrate leaderboard", err)
		panic(err)
	}
	log.Println("Leaderboard migrated successfully")

	return nil
}
