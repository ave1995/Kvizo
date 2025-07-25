package main

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

import (
	"kvizo-api/database"
	"kvizo-api/handlers"
	"kvizo-api/internal/auth"
	"kvizo-api/internal/loggers"
	"kvizo-api/internal/middlewares"
	"kvizo-api/services"
	"log/slog"
	"os"
	"time"

	"github.com/gin-gonic/gin"

	_ "kvizo-api/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	loggers.Init()

	gorm, err := database.NewDatabaseConnection()
	if err != nil {
		slog.Error("failed to ini Database Connection", slog.Any("error", err))
		os.Exit(1)
	}
	quizRepository, err := database.NewDatabaseQuizRepository(gorm)
	if err != nil {
		slog.Error("failed to ini Quiz Repository", slog.Any("error", err))
		os.Exit(1)
	}
	questionRepository, err := database.NewDatabaseQuestionRepository(gorm, quizRepository)
	if err != nil {
		slog.Error("failed to ini Question Repository", slog.Any("error", err))
		os.Exit(1)
	}
	authRepository, err := auth.NewDatabaseUserRepository(gorm)
	if err != nil {
		slog.Error("failed to ini Authentication Repository", slog.Any("error", err))
		os.Exit(1)
	}

	quizService := services.NewQuizService(quizRepository)
	questionService := services.NewQuestionService(questionRepository)

	quizHandler := handlers.NewQuizHandler(quizService)
	questionHandler := handlers.NewQuestionHandler(questionService)

	jwtManager := auth.NewJWTManager("RadekSmrdi", time.Minute)
	authService := auth.NewAuthService(authRepository, jwtManager)
	authHandler := auth.NewAuthHandler(authService)

	r := gin.Default()

	r.Use(middlewares.ErrorLoggingMiddleware())

	protected := r.Group("/api")
	protected.Use(auth.AuthMiddleware(jwtManager))

	// Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//Quizzes actions
	protected.POST("/quizzes", quizHandler.CreateQuizHandler)
	protected.GET("/quizzes", quizHandler.GetQuizzesHandler)
	protected.GET("/quiz/:id", quizHandler.GetQuizHandler)
	protected.PUT("/quiz/:id", quizHandler.UpdateQuizHandler)
	protected.DELETE("/quiz/:id", quizHandler.DeleteQuizHandler)

	//Questions actions
	protected.GET("/quizzes/:quiz_id/questions", questionHandler.GetQuestionsForQuizHandler)
	protected.POST("/quizzes/:quiz_id/questions", questionHandler.CreateQuestionHandler)
	protected.PUT("/question/:id", questionHandler.UpdateQuestionHandler)
	protected.DELETE("/question/:id", questionHandler.DeleteQuestionHandler)

	authGroup := r.Group("/auth")
	{
		authGroup.POST("/register", authHandler.RegisterUserHandler)
		authGroup.POST("/login", authHandler.LoginUserHandler)
	}

	err = r.Run(":8080")
	if err != nil {
		slog.Error("failed to start Gin server", slog.Any("error", err))
		os.Exit(1)
	}
}
