package main

import (
	"fmt"
	"log"
	"net"

	"github.com/KETAKOM/go-study-app/api/gss"
	"github.com/KETAKOM/go-study-app/application/handler"

	"github.com/KETAKOM/go-study-app/application/config"
	"google.golang.org/grpc"
)

func main() {
	log.Printf("grpcgss started...")
	listenPort, err := net.Listen("tcp", ":1111")

	if err != nil {
		log.Fatalln(err)
	}

	server := grpc.NewServer()

	config, err := config.LoadConfig()
	if err != nil {
		fmt.Println("LoadConfig Failed", err)
	}

	handle := &handler.GssHandler{Config: config}
	//実装したい実処理をserverに渡す
	gss.RegisterGssServer(server, handle)

	server.Serve(listenPort)
}
