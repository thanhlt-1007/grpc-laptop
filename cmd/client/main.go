package main

import (
	"context"
	"flag"
	"log"

	"grpc-laptop/go_protos/messages/cpu_message"
	"grpc-laptop/go_protos/messages/laptop_message"
	"grpc-laptop/go_protos/messages/memory_message"
	"grpc-laptop/go_protos/services/laptop_service"

	"github.com/google/uuid"
	"google.golang.org/grpc"
)

func createLaptops(laptopClient laptop_service.LaptopServiceClient) {
	createLaptop(
		laptopClient,
		10,
		&cpu_message.CPU{
			NumberCores: 10,
			MinGhz:      10,
			MaxGhz:      100,
		},
		&memory_message.Memory{},
	)
	createLaptop(
		laptopClient,
		12,
		&cpu_message.CPU{
			NumberCores: 20,
			MinGhz:      20,
			MaxGhz:      200,
		},
		&memory_message.Memory{},
	)
}

func createLaptop(laptopClient laptop_service.LaptopServiceClient, priceUsd float64, cpu *cpu_message.CPU, memory *memory_message.Memory) {
	laptop := &laptop_message.Laptop{
		Id:       uuid.New().String(),
		PriceUsd: priceUsd,
		Cpu:      cpu,
		Memory:   memory,
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
	createLaptops(laptopClient)
}
