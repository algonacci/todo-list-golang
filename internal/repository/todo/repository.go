package todo

import "github.com/algonacci/todo-list-golang/internal/model"

type Repository interface {
	GetTodoList() ([]model.Todo, error)
}
