package consts

import "errors"

var (
	ErrBodyRequired = errors.New("todo body is required")
	ErrNoTodosFound = errors.New("no todos found")
	ErrTodoNotFound = errors.New("todos not found")
)
