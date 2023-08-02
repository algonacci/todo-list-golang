package usecase

import "github.com/algonacci/todo-list-golang/internal/model"

type TodoUseCase interface {
	GetTodoList() ([]model.Todo, error)
}
