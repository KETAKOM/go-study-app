package repository

import "github.com/KETAKOM/go-sample-todo/application/domain/model"

type BookRepository interface {
	GetAll() ([]*model.Book, error)
}
