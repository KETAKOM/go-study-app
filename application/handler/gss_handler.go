package handler

import (
	"context"
	"fmt"

	pb "github.com/KETAKOM/go-study-app/api/gss"
	"github.com/KETAKOM/go-study-app/application/config"
	"github.com/KETAKOM/go-study-app/application/usecase"
)

type GssHandler struct {
	Config config.Config
	UC     usecase.GetGoogleSpreadSheetUsecase
}

func (gh *GssHandler) GetSpreadSheet(
	ctx context.Context,
	s *pb.GetSpreadSheetRequest,
) (*pb.GetSpreadSheetResponce, error) {
	acuc := usecase.NewGetAccountCredentials(gh.Config)
	client, err := acuc.GetAccountCredentials()
	if err != nil {
		fmt.Println("GetAccountCredentials Failed", err)
	}
	table, err := acuc.ReadTable(client)
	if err != nil {
		fmt.Println("ReadTable Failed", err)
	}

	schema, err := usecase.Convater(table)
	if err != nil {
		fmt.Println("Convater Failed", err)
	}

	return &pb.GetSpreadSheetResponce{
		Res:    "OK",
		Schema: schema,
	}, nil
}
