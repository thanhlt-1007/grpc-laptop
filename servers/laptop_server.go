package servers

import (
	"context"
	"errors"
	"log"

	"grpc-laptop/go_protos/services/laptop_service"
	"grpc-laptop/stores"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type LaptopServer struct {
	store *stores.InMemoryLaptopStore
}

func NewLaptopServer(store *stores.InMemoryLaptopStore) *LaptopServer {
	server := &LaptopServer{
		store: store,
	}
	return server
}

func (server *LaptopServer) CreateLaptop(
	ctx context.Context,
	req *laptop_service.CreateLaptopRequest,
) (*laptop_service.CreateLaptopResponse, error) {
	laptop := req.GetLaptop()
	log.Printf("LaptopServer CreateLaptop() req.GetLaptop(): %#v\n", laptop)

	id := laptop.Id
	if len(id) > 0 {
		_, err := uuid.Parse(id)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "Laptop ID is not a valid UUID: %v", err)
		}
	} else {
		id, err := uuid.NewRandom()
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Cannot generate a new laptio ID: %v", err)
		}

		laptop.Id = id.String()
	}

	err := server.store.Save(laptop)
	if err != nil {
		code := codes.Internal
		if errors.Is(err, stores.ErrAlreadyExists) {
			code = codes.AlreadyExists
		}

		return nil, status.Errorf(code, "can't save laptop to the store: %#v", err)
	}

	log.Printf("saved laptop with id: %s", laptop.Id)
	response := laptop_service.CreateLaptopResponse{
		Id: laptop.Id,
	}
	return &response, nil
}
