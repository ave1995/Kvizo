package main

import (
	"kvizo-api/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDatabase()

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	r.GET("/testdb", func(c *gin.Context) {
		sqlDB, err := database.DB.DB()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		err = sqlDB.Ping()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database not reachable"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Database connection successful"})
	})

	r.Run(":8080")
}
