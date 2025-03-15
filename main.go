package main

import (
	"github/pull_request_webhook/infrastructure"
	"github/workflow/infraestructure"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	engine := gin.Default()

	infrastructure.Routes(engine)
	infraestructure.Routes(engine)

	engine.Run(":" + port)

}
