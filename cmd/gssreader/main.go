package main

import (
	"fmt"
	"os"

	"github.com/KETAKOM/go-study-app/application/config"
	"github.com/KETAKOM/go-study-app/application/domain/model"
	uc "github.com/KETAKOM/go-study-app/application/usecase"
	"github.com/yokoe/herschel"
	"github.com/yokoe/herschel/option"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		fmt.Println("LoadConfig Failed", err)
	}
	fmt.Println(config)

	dir := os.Getenv("JSON_DIR")
	op := option.WithServiceAccountCredentials(dir)
	client, err := herschel.NewClient(op)
	if err != nil {
		fmt.Println(err)
		return
	}
	id := os.Getenv("SHEET_ID")
	name := os.Getenv("SHEET_NAME")

	table, err := client.ReadTable(id, name)
	if err != nil {
		fmt.Println(err)
		return
	}

	usecase := uc.NewCreateSQLUseCase()
	usecase.SchemaRepository = &model.Schema{}
	usecase.SQLRepository = &model.SQL{}

	var sql = usecase.CreateSQL(table)

	fmt.Println(sql.Query)
}
