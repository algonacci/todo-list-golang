package todo

import (
	"github.com/algonacci/todo-list-golang/internal/model"
	"gorm.io/gorm"
)

type todoRepo struct {
	db *gorm.DB
}

func GetRepository(db *gorm.DB) Repository {
	return &todoRepo{
		db: db,
	}
}

func (t *todoRepo) GetTodoList() ([]model.Todo, error) {
	var todos []model.Todo
	result := t.db.Find(&todos)
	if result.Error != nil {
		return nil, result.Error
	}
	return todos, nil
}
