package handler

import (
	"context"
	"fmt"

	pb "github.com/KETAKOM/go-study-app/api/gss"
	"github.com/KETAKOM/go-study-app/application/config"
	"github.com/KETAKOM/go-study-app/application/domain/model"
	"github.com/KETAKOM/go-study-app/application/usecase"
)

type GssHandler struct {
	Config config.Config
}

func (gh *GssHandler) GetQueryBySpreadSheet(
	ctx context.Context,
	s *pb.GetQueryBySpreadSheetRequest,
) (*pb.GetQueryBySpreadSheetResponce, error) {
	ucac := usecase.NewGetAccountCredentials(gh.Config)
	client, err := ucac.GetAccountCredentials()
	if err != nil {
		fmt.Println("GetAccountCredentials Failed", err)
	}
	table, err := ucac.ReadTable(client)
	if err != nil {
		fmt.Println("ReadTable Failed", err)
	}

	queryUsecase := usecase.NewCreateSQLUseCase()
	queryUsecase.SchemaRepository = &model.Schema{}
	queryUsecase.SQLRepository = &model.SQL{}

	sql, err := queryUsecase.CreateSQL(table)
	if err != nil {
		return &pb.GetQueryBySpreadSheetResponce{
			Res:   "Failed",
			Query: "",
		}, err
	}

	return &pb.GetQueryBySpreadSheetResponce{
		Res:   "OK",
		Query: sql.Query,
	}, nil
}
