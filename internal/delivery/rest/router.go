package rest

import (
	"github.com/algonacci/todo-list-golang/internal/usecase"
	"github.com/labstack/echo/v4"
)

func LoadRoutes(e *echo.Echo, todoUseCase usecase.TodoUseCase) {
	todoHandler := NewTodoHandler(todoUseCase)
	e.GET("/", todoHandler.GetHelloWorld)
	e.GET("/todos", todoHandler.GetTodoList)
}
