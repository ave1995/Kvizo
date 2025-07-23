package database

import (
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabaseConnection() (*gorm.DB, error) {
	godotenv.Load()
	viper.AutomaticEnv()

	host := viper.GetString("DB_HOST")
	user := viper.GetString("DB_USER")
	password := viper.GetString("DB_PASSWORD")
	dbname := viper.GetString("DB_NAME")
	port := viper.GetString("DB_PORT")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		host, user, password, dbname, port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	err = db.AutoMigrate(&databaseQuiz{}, &databaseQuestion{})
	if err != nil {
		log.Fatalf("AutoMigrate failed: %v", err)
	}

	return db, nil
}

func getByID[T any](tx *gorm.DB, id uuid.UUID) (*T, error) {
	var model T
	if err := tx.First(&model, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &model, nil
}
