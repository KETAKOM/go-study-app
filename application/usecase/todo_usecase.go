package usecase

import (
	"github.com/KETAKOM/go-study-app/application/domain/model"
	"github.com/KETAKOM/go-study-app/application/domain/repository"
	pre "github.com/KETAKOM/go-study-app/application/infra/persistence"
)

type TodoUsecase struct{}

func (todo TodoUsecase) GetList() ([]*model.Todo, error) {
	var tl []*model.Todo
	var err error
	tl, err = repository.TodoRepository(pre.TodoPersistence{}).GetList()

	if err != nil {
		return nil, err
	}
	return tl, nil
}

func (todo TodoUsecase) AddTodo(t *model.Todo) (bool, error) {
	var err error
	_, err = repository.TodoRepository(pre.TodoPersistence{}).AddTodo(t)
	if err != nil {
		return false, err
	}
	return true, nil
}
