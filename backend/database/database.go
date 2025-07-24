package database

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConfig struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
}

func loadDBConfig() (*DBConfig, error) {
	godotenv.Load()
	viper.AutomaticEnv()

	var cfg DBConfig

	cfg.Host = viper.GetString("DB_HOST")
	cfg.User = viper.GetString("DB_USER")
	cfg.Password = viper.GetString("DB_PASSWORD")
	cfg.DBName = viper.GetString("DB_NAME")
	cfg.Port = viper.GetString("DB_PORT")

	return &cfg, nil
}

func NewDatabaseConnection() (*gorm.DB, error) {
	cfg, err := loadDBConfig()
	if err != nil {
		return nil, fmt.Errorf("loadDBConfig: %w", err)
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		cfg.Host, cfg.User, cfg.Password, cfg.DBName, cfg.Port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("gorm.Open: %w", err)
	}

	err = db.AutoMigrate(&databaseQuiz{}, &databaseQuestion{})
	if err != nil {
		return nil, fmt.Errorf("db.AutoMigrate: %w", err)
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

type uuidLike interface {
	~[16]byte
}

func toUUID[T uuidLike](id T) uuid.UUID {
	return uuid.UUID(id)
}
