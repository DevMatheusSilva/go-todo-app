package handlers

import (
	"fmt"
	"github.com/DevMatheusSilva/go-todo-app/cmd/internal/domain"
	"github.com/DevMatheusSilva/go-todo-app/cmd/internal/utils/consts"
	"github.com/gofiber/fiber/v2"
)

var todos []domain.Todo

func CreateTodo(c *fiber.Ctx) error {
	todo := &domain.Todo{}

	if err := c.BodyParser(todo); err != nil {
		return err
	}

	if todo.Body == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": consts.ErrBodyRequired.Error()})
	}

	todo.ID = len(todos) + 1
	todos = append(todos, *todo)

	return c.Status(fiber.StatusCreated).JSON(todo)
}

func ShowAllTodos(c *fiber.Ctx) error {
	if todos == nil || len(todos) == 0 {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": consts.ErrNoTodosFound.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(todos)
}

func CompleteTodo(c *fiber.Ctx) error {
	id := c.Params("id")

	for i, todo := range todos {
		if fmt.Sprint(todo.ID) == id {
			todos[i].Completed = !todos[i].Completed
			return c.Status(fiber.StatusOK).JSON(todos[i])
		}
	}

	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": consts.ErrTodoNotFound.Error()})
}

func DeleteTodo(c *fiber.Ctx) error {
	id := c.Params("id")

	for i, todo := range todos {
		if fmt.Sprint(todo.ID) == id {
			todos = append(todos[:i], todos[i+1:]...)
			return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": consts.DeletedSuccessfullyMsg})
		}
	}

	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": consts.ErrTodoNotFound.Error()})
}
