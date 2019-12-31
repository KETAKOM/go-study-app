package main

import (
	"log"
	"net"

	cque "github.com/KETAKOM/go-study-app/api/cque"
	"github.com/KETAKOM/go-study-app/application/handler"
	"google.golang.org/grpc"
)

func main() {
	log.Printf("cque started...")
	listenPort, err := net.Listen("tcp", ":1122")

	if err != nil {
		log.Fatalln(err)
	}

	server := grpc.NewServer()

	handle := &handler.FileHandler{}
	//実装したい実処理をserverに渡す
	cque.RegisterCuqeServer(server, handle)

	server.Serve(listenPort)

}
