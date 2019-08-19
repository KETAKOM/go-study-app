package main

import (
	"log"
	"net"

	"github.com/KETAKOM/go-study-app/api/hello"
	"github.com/KETAKOM/go-study-app/application/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	log.Printf("grpchello started...")
	listenPort, err := net.Listen("tcp", ":1234")

	if err != nil {
		log.Fatalln(err)
	}

	server := grpc.NewServer()
	reflection.Register(server)
	helloService := &service.HelloService{}
	//実装したい実処理をserverに渡す
	hello.RegisterHelloServer(server, helloService)

	server.Serve(listenPort)
}
