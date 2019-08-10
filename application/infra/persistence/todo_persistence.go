package persistence

import (
	"time"

	"github.com/KETAKOM/go-study-app/application/domain/model"
	_ "github.com/go-sql-driver/mysql"
)

type TodoPersistence struct{}

func (todo TodoPersistence) GetList() ([]*model.Todo, error) {
	t := model.Todo{}
	t.Id = 1
	t.Title = "Golangの勉強"
	t.Detail = "gRPC使ってみる"
	t.Auther = "山田 太郎"
	t.CreatedAt = time.Now()
	t.UpdatedAt = time.Now()
	return []*model.Todo{&t}, nil
}
