package clients_test

import (
	"context"
	"net"
	"testing"

	"grpc-laptop/go_protos/messages/laptop_message"
	"grpc-laptop/go_protos/services/laptop_service"
	"grpc-laptop/servers"
	"grpc-laptop/stores"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
)

func TestLaptopClientCreateLaptop(test *testing.T) {
	test.Parallel()

	_, serverAddr := startLaptopServerTest(test)
	laptopClient := newLaptopClientTest(test, serverAddr)

	laptop := &laptop_message.Laptop{
		Id: uuid.New().String(),
	}

	request := &laptop_service.CreateLaptopRequest{
		Laptop: laptop,
	}

	response, err := laptopClient.CreateLaptop(context.Background(), request)
	require.NoError(test, err)
	require.NotNil(test, response)
	require.Equal(test, response.Id, laptop.Id)
}

func startLaptopServerTest(test *testing.T) (*servers.LaptopServer, string) {
	laptopServer := servers.NewLaptopServer(stores.NewInMemoryLaptopStore())

	grpcServer := grpc.NewServer()
	laptop_service.RegisterLaptopServiceServer(grpcServer, laptopServer)

	listener, err := net.Listen("tcp", ":0")
	require.NoError(test, err)

	go grpcServer.Serve(listener)

	return laptopServer, listener.Addr().String()
}

func newLaptopClientTest(test *testing.T, serverAddr string) laptop_service.LaptopServiceClient {
	grpcClientConn, err := grpc.Dial(serverAddr, grpc.WithInsecure())
	require.NoError(test, err)

	client := laptop_service.NewLaptopServiceClient(grpcClientConn)
	return client
}
