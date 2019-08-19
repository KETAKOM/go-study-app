package service

import (
	"context"

	pb "github.com/KETAKOM/go-study-app/api/hello"
)

type HelloService struct {
}

//作られたxxx.pb.goのinterfaceを満たすように実装する。

func (hs *HelloService) GetMessge(ctx context.Context, message *pb.MessageRequest) (*pb.MessageResponce, error) {
	return &pb.MessageResponce{
		Res: "sample-message",
	}, nil
}
