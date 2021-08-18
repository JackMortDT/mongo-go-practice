package app

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print(".env file found D:")
	}
}

var (
	router = gin.Default()
)

func StartApp() {
	mongoUrl := os.Getenv("MONGO_URL")

	startDatabase(mongoUrl)
	fmt.Println("DATABASE STARTED")

	routes()
	router.Run(":8080")
}
