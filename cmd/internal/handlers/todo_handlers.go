package handlers

import (
	"github.com/DevMatheusSilva/go-todo-app/cmd/internal/config/db"
	"github.com/DevMatheusSilva/go-todo-app/cmd/internal/domain"
	"github.com/DevMatheusSilva/go-todo-app/cmd/internal/repository"
	"github.com/DevMatheusSilva/go-todo-app/cmd/internal/utils/consts"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateTodo(database *db.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		todo := new(domain.Todo)
		if err := c.BodyParser(todo); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}
		if todo.Body == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": consts.ErrBodyRequired})
		}

		var savedTodo, err = repository.SaveTodo(c.Context(), database, todo)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		return c.Status(fiber.StatusCreated).JSON(savedTodo)
	}
}

func ShowAllTodos(database *db.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var todos, err = repository.GetAllTodos(c.Context(), database)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		return c.Status(fiber.StatusOK).JSON(todos)
	}
}

func CompleteTodo(database *db.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		objectId, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": consts.ErrInvalidID})
		}

		err = repository.CompleteTodo(c.Context(), database, objectId)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": consts.UpdatedSuccessfullyMsg})
	}
}

func DeleteTodo(database *db.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		objectId, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": consts.ErrInvalidID})
		}

		err = repository.DeleteTodo(c.Context(), database, objectId)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": consts.DeletedSuccessfullyMsg})
	}
}
