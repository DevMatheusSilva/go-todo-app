package routes

import (
	"github.com/DevMatheusSilva/go-todo-app/cmd/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupTodoRoutes(app *fiber.App) {
	app.Post("/api/todos", handlers.CreateTodo)
	app.Get("/api/todos", handlers.ShowAllTodos)
	app.Patch("/api/todos/:id", handlers.CompleteTodo)
	app.Delete("/api/todos/:id", handlers.DeleteTodo)
}
