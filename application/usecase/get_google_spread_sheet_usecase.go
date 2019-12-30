package usecase

import (
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
