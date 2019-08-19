package service

import (
	"context"
	pb "sample/api/hello"
)

type HelloService struct {
}

//作られたxxx.pb.goのinterfaceを満たすように実装する。

func (hs *HelloService) GetMessge(ctx context.Context, message *pb.MessageRequest) (*pb.MessageResponce, error) {
	return &pb.MessageResponce{
		Res: "sample-message",
	}, nil
}
