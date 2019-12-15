package main

import (
	"fmt"

	pb "github.com/KETAKOM/go-study-app/api/sample"
	"google.golang.org/grpc"

	"context"
)

func main() {
	address := "localhost:2222"
	// Set up a connection to the server.
	ctx := context.Background()
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	c := pb.NewSampleClient(conn)

	r, err := c.GetSample(ctx, &pb.SampleRequest{Sample: "ikerukana"})
	if err != nil {
		fmt.Println("error", err)
		return
	}

	fmt.Println(r)
}
