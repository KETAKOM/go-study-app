package usecase

import (
	"github.com/KETAKOM/go-sample-todo/application/domain/model"
	brep "github.com/KETAKOM/go-sample-todo/application/domain/repository"
	preb "github.com/KETAKOM/go-sample-todo/application/infra/persistence"
)

type BookUsecase struct{}

func (bu BookUsecase) GetAll() ([]*model.Book, error) {
	books, err := brep.BookRepository(preb.BookPersistence{}).GetAll()

	if err != nil {
		return nil, err
	}
	return books, nil
}
