package main

import (
	"log"
	"net"

	"github.com/KETAKOM/go-study-app/api/sample"
	"github.com/KETAKOM/go-study-app/application/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	log.Printf("grpcSample started...")
	listenPort, err := net.Listen("tcp", ":2222")

	if err != nil {
		log.Fatalln(err)
	}

	server := grpc.NewServer()
	reflection.Register(server)
	sampleService := &service.SampleService{}
	//実装したい実処理をserverに渡す
	sample.RegisterSampleServer(server, sampleService)

	server.Serve(listenPort)
}
