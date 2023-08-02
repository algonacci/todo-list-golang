package usecase

import (
	"github.com/algonacci/todo-list-golang/internal/model"
	"github.com/algonacci/todo-list-golang/internal/repository/todo"
)

type todoUseCaseImpl struct {
	todoRepo todo.Repository
}

func NewTodoUseCase(todoRepo todo.Repository) TodoUseCase {
	return &todoUseCaseImpl{
		todoRepo: todoRepo,
	}
}

func (uc *todoUseCaseImpl) GetTodoList() ([]model.Todo, error) {
	return uc.todoRepo.GetTodoList()
}
