package repository

import "github.com/KETAKOM/go-study-app/application/domain/model"

type BookRepository interface {
	GetAll() ([]*model.Book, error)
}
