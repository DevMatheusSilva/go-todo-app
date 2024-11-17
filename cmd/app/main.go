package main

import (
	"github.com/DevMatheusSilva/go-todo-app/cmd/internal/routes"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := fiber.New()
	routes.SetupTodoRoutes(app)
	log.Fatal(app.Listen(":3000"))
}
