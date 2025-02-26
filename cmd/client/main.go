package main

import (
	"context"
	"flag"
	"log"

	"grpc-laptop/go_protos/messages/laptop_message"
	"grpc-laptop/go_protos/services/laptop_service"

	"github.com/google/uuid"
	"google.golang.org/grpc"
)

func createLaptop(laptopClient laptop_service.LaptopServiceClient) {
	laptop := &laptop_message.Laptop{
		Id: uuid.New().String(),
	}

	request := &laptop_service.CreateLaptopRequest{
		Laptop: laptop,
	}

	_, err := laptopClient.CreateLaptop(context.Background(), request)
	if err != nil {
		log.Fatalf("can't create laptop: %#v", err)
		return
	}

	log.Printf("created laptop with Id: %s\n", laptop.Id)
}

func main() {
	address := flag.String("address", "", "server address")
	flag.Parse()
	log.Printf("Run client on address %s\n", *address)

	conn, err := grpc.Dial(*address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("can't dial server: %#v", err)
	}

	laptopClient := laptop_service.NewLaptopServiceClient(conn)
	createLaptop(laptopClient)
}
