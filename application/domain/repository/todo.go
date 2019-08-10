package repository

import "github.com/KETAKOM/go-study-app/application/domain/model"

type TodoRepository interface {
	GetList() ([]*model.Todo, error)
}