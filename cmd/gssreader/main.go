package main

import (
	"fmt"

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

	op := option.WithServiceAccountCredentials(config.GoogleSpreadSheet.AUTH_DIR)
	client, err := herschel.NewClient(op)
	if err != nil {
		fmt.Println(err)
		return
	}

	table, err := client.ReadTable(
		config.GoogleSpreadSheet.SHEET_ID,
		config.GoogleSpreadSheet.SHEET_NAME,
	)
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
