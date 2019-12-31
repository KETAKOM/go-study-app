package handler

import (
	"context"

	pb "github.com/KETAKOM/go-study-app/api/cque"
	uc "github.com/KETAKOM/go-study-app/application/usecase"
)

type FileHandler struct{}

func (ch *FileHandler) GetFile(
	ctx context.Context,
	s *pb.GetFileRequest,
) (*pb.GetFileResponce, error) {
	usecase := uc.NewCreateFileUseCase()
	err := usecase.CreateFile(s.Query)
	if err != nil {
		return &pb.GetFileResponce{
			Result: "Failed",
		}, err
	}

	return &pb.GetFileResponce{
		Result: "OK",
	}, nil
}
