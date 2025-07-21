package main

import (
	"kvizo-api/database"
	"kvizo-api/handlers"

	"github.com/gin-gonic/gin"

	_ "kvizo-api/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	database.ConnectDatabase()

	r := gin.Default()

	// Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//Test actions
	r.GET("/ping", handlers.PingHandler)
	r.GET("/testdb", handlers.TestDbHandler)

	//Quizzes actions
	r.POST("/quizzes", handlers.CreateQuizHandler)
	r.GET("/quizzes", handlers.GetQuizzesHandler)
	r.GET("/quiz/:id", handlers.GetQuizHandler)

	r.Run(":8080")
}
