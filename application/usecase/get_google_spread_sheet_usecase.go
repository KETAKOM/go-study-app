package usecase

import (
	"fmt"

	pb "github.com/KETAKOM/go-study-app/api/gss"
	"github.com/KETAKOM/go-study-app/application/config"
	"github.com/yokoe/herschel"
	"github.com/yokoe/herschel/option"
)

type GetGoogleSpreadSheetUsecase struct {
	conf config.GoogleSpreadSheetConfig
}

func NewGetAccountCredentials(conf config.Config) *GetGoogleSpreadSheetUsecase {
	return &GetGoogleSpreadSheetUsecase{conf.GoogleSpreadSheet}
}

func (g *GetGoogleSpreadSheetUsecase) ReadTable(client *herschel.Client) (*herschel.Table, error) {

	table, err := client.ReadTable(
		g.conf.SHEET_ID,
		g.conf.SHEET_NAME,
	)

	if err != nil {
		return nil, err
	}

	return table, nil
}

func (g *GetGoogleSpreadSheetUsecase) GetAccountCredentials() (*herschel.Client, error) {
	op := option.WithServiceAccountCredentials(g.conf.AUTH_DIR)
	client, err := herschel.NewClient(op)

	if err != nil {
		return nil, err
	}
	return client, err
}

func Convater(t *herschel.Table) (*pb.GetSpreadSheetResponce_Schema, error) {

	var r *pb.GetSpreadSheetResponce_Schema

	rows := t.GetRows()
	fmt.Println(r.Columns)

	for i := 2; i < rows; i++ {

		// r.Columns[int32(i)].Id = t.GetStringValue(i, 0)
		// fmt.Println(r.Columns[int32(i)].Id)
		// r.Columns[int32(i)].LogicalName = t.GetStringValue(i, 1)
		// r.Columns[int32(i)].PhysicalName = t.GetStringValue(i, 2)
		// r.Columns[int32(i)].Type = t.GetStringValue(i, 3)
		// r.Columns[int32(i)].NullAble = t.GetStringValue(i, 4)
		// r.Columns[int32(i)].PK = t.GetStringValue(i, 5)
		// r.Columns[int32(i)].Default = t.GetStringValue(i, 6)
	}

	return r, nil
}
