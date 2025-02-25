package main

import (
	"flag"
	"fmt"
	"grpc-laptop/go_protos/services/laptop_service"
	"grpc-laptop/servers"
	"grpc-laptop/stores"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	port := flag.Int("port", 0, "server port")
	flag.Parse()
	log.Printf("Run server on port %d\n", *port)

	store := stores.NewInMemoryLaptopStore()
	server := servers.NewLaptopServer(store)
	grpcServer := grpc.NewServer()

	laptop_service.RegisterLaptopServiceServer(grpcServer, server)

	address := fmt.Sprintf("0.0.0.0:%d", *port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("can't start server: %#v", err)
	}

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("can't start server: %#v", err)
	}
}
