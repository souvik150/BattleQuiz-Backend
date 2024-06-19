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

	err = DB.AutoMigrate(&models.User{}, &models.MCQ{}).Error
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
		return err
	}
	log.Println("Database migrated successfully")

	return nil
}
