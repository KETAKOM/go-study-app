package persistence

import (
	"database/sql"

	"github.com/KETAKOM/go-study-app/application/domain/model"
	"github.com/KETAKOM/go-study-app/application/lib"
	_ "github.com/go-sql-driver/mysql"
)

type TodoPersistence struct {
	Connect *sql.DB
}

func (todo TodoPersistence) GetList() ([]*model.Todo, error) {

	db, err := lib.NewDBConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("select id, title, detail, auther from todo")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// 構造体の型が入る配列を定義
	ts := make([]*model.Todo, 0)

	// レコード一件一件を用意した構造体に当てはめていく
	for rows.Next() {
		t := model.Todo{}
		err := rows.Scan(&t.Id, &t.Title, &t.Detail, &t.Auther)
		if err != nil {
			return nil, err
		}
		ts = append(ts, &t)
	}
	return ts, nil
}
