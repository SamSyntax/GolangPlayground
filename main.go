package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var apiKey string

func main() {
	apiKey = os.Getenv("SECRET")
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()

	r.Use(AuthMiddleWare())

	r.GET("/data", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Data fetched"})
		c.JSON(http.StatusOK, gin.H{"message": apiKey})

	})

	r.Run(":3030")

	print(apiKey)

}
func AuthMiddleWare() gin.HandlerFunc {
	apiKey = os.Getenv("SECRET")

	return func(c *gin.Context) {
		clientAPIKey := c.GetHeader("X-API-Key")

		if clientAPIKey != apiKey {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized Access"})
			c.JSON(http.StatusOK, gin.H{"message": apiKey})
			c.Abort()

			return
		}

		c.Next()
	}
}
