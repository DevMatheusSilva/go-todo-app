package routes

import (
	"github.com/DevMatheusSilva/go-todo-app/cmd/internal/config/db"
	"github.com/DevMatheusSilva/go-todo-app/cmd/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupTodoRoutes(app *fiber.App, database *db.Database) {
	app.Post("/api/todos", handlers.CreateTodo(database))
	app.Get("/api/todos", handlers.ShowAllTodos(database))
	app.Patch("/api/todos/:id", handlers.CompleteTodo(database))
	app.Delete("/api/todos/:id", handlers.DeleteTodo(database))
}
