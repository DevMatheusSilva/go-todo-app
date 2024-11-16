package main

import (
	"context"
	"fmt"
	"github.com/DevMatheusSilva/go-todo-app/cmd/internal/config/db"
	"github.com/DevMatheusSilva/go-todo-app/cmd/internal/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load("../internal/config/.env")
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	app := fiber.New()
	ctx := context.Background()

	client, database, err := db.Connect(ctx)
	if err != nil {
		fmt.Println("Error connecting to th database:", err)
		return
	}
	defer client.Disconnect(ctx)

	routes.SetupTodoRoutes(app, database)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Fatal(app.Listen("0.0.0.0:" + port))
}
