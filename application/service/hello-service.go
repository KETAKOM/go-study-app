package service

import (
	"context"

	"google.golang.org/grpc"

	pb "github.com/KETAKOM/go-study-app/api/hello"
	sp "github.com/KETAKOM/go-study-app/api/sample"
)

type HelloService struct {
}

//作られたxxx.pb.goのinterfaceを満たすように実装する。

func (hs *HelloService) GetMessge(ctx context.Context, message *pb.MessageRequest) (*pb.MessageResponce, error) {

	if message.Message == "" {
		return nil, nil
	}

	address := "localhost:2222"
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	c := sp.NewSampleClient(conn)

	r, err := c.GetSample(ctx, &sp.SampleRequest{Sample: "ikerukana"})
	if err != nil {
		return nil, err
	}

	return &pb.MessageResponce{
		Res: "sample :" + message.Message + ": " + r.Res,
	}, nil
}
