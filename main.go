package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/omerfruk/Go-generics-api/database"
	"github.com/omerfruk/Go-generics-api/router"
)

func main() {
	_ = godotenv.Load()

	if err := database.DBConnect(); err != nil {
		log.Fatalf("connect to database: %v", err)
	}
	if err := database.AutoMigrate(); err != nil {
		log.Fatalf("run database migrations: %v", err)
	}

	app := fiber.New()
	router.Setup(app)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "4747"
	}

	log.Fatal(app.Listen(":" + port))
}
