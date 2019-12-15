package service

import (
	"context"

	pb "github.com/KETAKOM/go-study-app/api/sample"
)

type SampleService struct{}

func (s *SampleService) GetSample(ctx context.Context, sample *pb.SampleRequest) (*pb.SampleResponce, error) {
	return &pb.SampleResponce{
		Res: "ikeruyo",
	}, nil
}
