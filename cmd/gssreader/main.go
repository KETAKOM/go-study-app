package main

import (
	"fmt"

	"github.com/KETAKOM/go-study-app/application/config"
	"github.com/KETAKOM/go-study-app/application/domain/model"
	uc "github.com/KETAKOM/go-study-app/application/usecase"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		fmt.Println("LoadConfig Failed", err)
	}

	acuc := uc.NewGetAccountCredentials(config)
	client, err := acuc.GetAccountCredentials()
	if err != nil {
		fmt.Println("GetccountCredentials Failed", err)
	}
	table, err := acuc.ReadTable(client)
	if err != nil {
		fmt.Println("ReadTable Failed", err)
	}

	usecase := uc.NewCreateSQLUseCase()
	usecase.SchemaRepository = &model.Schema{}
	usecase.SQLRepository = &model.SQL{}

	var sql = usecase.CreateSQL(table)

	fmt.Println(sql.Query)
}
