package main

import (
	"kvizo-api/database"
	"kvizo-api/handlers"
	"kvizo-api/services"

	"github.com/gin-gonic/gin"

	_ "kvizo-api/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	gorm, _ := database.NewDatabaseConnection()
	quizRepository := database.NewDatabaseQuizRepository(gorm)
	questionRepository := database.NewDatabaseQuestionRepository(gorm, quizRepository)

	quizService := services.NewQuizService(quizRepository)
	questionService := services.NewQuestionService(questionRepository)

	quizHandler := handlers.NewQuizHandler(quizService)
	questionHandler := handlers.NewQuestionHandler(questionService)

	r := gin.Default()

	//TODO: přidej verzování v1 groups
	//TODO: správná práce s context
	//TODO: slog logy

	// Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//Test actions
	r.GET("/ping", handlers.PingHandler)

	//Quizzes actions
	r.POST("/quizzes", quizHandler.CreateQuizHandler)
	r.GET("/quizzes", quizHandler.GetQuizzesHandler)
	r.GET("/quiz/:id", quizHandler.GetQuizHandler)

	//Questions actions
	r.GET("/quizzes/:quiz_id/questions", questionHandler.GetQuestionsForQuizHandler)
	r.POST("/quizzes/:quiz_id/questions", questionHandler.CreateQuestionHandler)

	r.Run(":8080")
}
