package main

import (
	"fmt"

	"context"

	pbc "github.com/KETAKOM/go-study-app/api/cque"
	pbg "github.com/KETAKOM/go-study-app/api/gss"
	"google.golang.org/grpc"
)

func main() {
	address := "localhost:1111"
	// Set up a connection to the server.
	ctx := context.Background()
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	c := pbg.NewGssClient(conn)

	s := pbg.GetQueryBySpreadSheetRequest{}
	r, err := c.GetQueryBySpreadSheet(ctx, &s)
	if err != nil {
		fmt.Println("error GetSpreadSheet: ", err)
		return
	}

	address = "localhost:1122"
	// Set up a connection to the server.
	ctx = context.Background()
	conn, err = grpc.Dial(address, grpc.WithInsecure())
	cc := pbc.NewCuqeClient(conn)

	ss := pbc.GetFileRequest{
		Query: r.Query,
	}
	_, err = cc.GetFile(ctx, &ss)
	if err != nil {
		fmt.Println("error GetSpreadSheet: ", err)
		return
	}
}
