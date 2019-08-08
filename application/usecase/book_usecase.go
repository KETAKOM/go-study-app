package usecase

import (
	"github.com/KETAKOM/go-study-app/application/domain/model"
	brep "github.com/KETAKOM/go-study-app/application/domain/repository"
	preb "github.com/KETAKOM/go-study-app/application/infra/persistence"
)

type BookUsecase struct{}

func (bu BookUsecase) GetAll() ([]*model.Book, error) {
	var books []*model.Book
	var err error

	books, err = brep.BookRepository(preb.BookPersistence{}).GetAll()

	if err != nil {
		return nil, err
	}
	return books, nil
}
