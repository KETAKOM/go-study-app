package usecase

import (
	"fmt"
	"os"
	"time"
)

type CreateFileUseCase struct{}

func NewCreateFileUseCase() *CreateSQLUseCase {
	return &CreateSQLUseCase{}
}

func (uc *CreateSQLUseCase) CreateFile(query string) error {
	t := time.Now()
	fileName := "query" + t.Format("20060102150405") + ".sql"
	dir := "./"
	path := dir + fileName
	fp, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
	}
	defer fp.Close()

	fmt.Fprintln(fp, query)
	return nil
}
