package main

import (
	"fmt"

	pb "github.com/KETAKOM/go-study-app/api/hello"
	"google.golang.org/grpc"

	"context"
)

func main() {
	address := "localhost:1234"
	// Set up a connection to the server.
	ctx := context.Background()
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	c := pb.NewHelloClient(conn)

	r, err := c.GetMessge(ctx, &pb.MessageRequest{Message: "ikerukana"})
	if err != nil {
		fmt.Println("error", err)
		return
	}

	fmt.Println(r)
}
