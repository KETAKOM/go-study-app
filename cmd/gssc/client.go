package main

import (
	"fmt"

	pb "github.com/KETAKOM/go-study-app/api/gss"
	"google.golang.org/grpc"

	"context"
)

func main() {
	address := "localhost:1111"
	// Set up a connection to the server.
	ctx := context.Background()
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	c := pb.NewGssClient(conn)

	s := pb.GetSpreadSheetRequest{}
	r, err := c.GetSpreadSheet(ctx, &s)
	if err != nil {
		fmt.Println("error GetSpreadSheet: ", err)
		return
	}

	fmt.Println(r.Res)
}
