package main

import (
	"context"
	"fmt"

	pb "github.com/KETAKOM/go-study-app/api/cque"
	"google.golang.org/grpc"
)

func main() {
	address := "localhost:1122"
	// Set up a connection to the server.
	ctx := context.Background()
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	c := pb.NewCuqeClient(conn)

	s := pb.GetFileRequest{
		Query: "aiueo cque",
	}
	_, err = c.GetFile(ctx, &s)
	if err != nil {
		fmt.Println("error GetSpreadSheet: ", err)
		return
	}
}
