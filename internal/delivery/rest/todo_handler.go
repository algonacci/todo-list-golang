package rest

import (
	"net/http"

	"github.com/algonacci/todo-list-golang/internal/usecase"
	"github.com/labstack/echo/v4"
)

type TodoHandler struct {
	todoUseCase usecase.TodoUseCase
}

func NewTodoHandler(todoUseCase usecase.TodoUseCase) *TodoHandler {
	return &TodoHandler{
		todoUseCase: todoUseCase,
	}
}

func (h *TodoHandler) GetTodoList(c echo.Context) error {
	todos, err := h.todoUseCase.GetTodoList()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, todos)
}
