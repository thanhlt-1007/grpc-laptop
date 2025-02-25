package servers_test

import (
	"context"
	"grpc-laptop/go_protos/messages/laptop_message"
	"grpc-laptop/go_protos/services/laptop_service"
	"grpc-laptop/servers"
	"grpc-laptop/stores"
	"testing"

	"github.com/google/uuid"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TestNewLaptopServerTestCase struct {
	name   string
	laptop *laptop_message.Laptop
	store  *stores.InMemoryLaptopStore
	code   codes.Code
}

func TestNewLaptopServer(test *testing.T) {
	test.Parallel()

	newLaptopWithId := &laptop_message.Laptop{
		Id: uuid.New().String(),
	}

	newLaptopWithoutId := &laptop_message.Laptop{
		Id: "",
	}

	newLaptopWithInvalidId := &laptop_message.Laptop{
		Id: "invalid",
	}

	savedLaptop := &laptop_message.Laptop{
		Id: uuid.New().String(),
	}
	savedStore := stores.NewInMemoryLaptopStore()
	err := savedStore.Save(savedLaptop)
	require.Nil(test, err)
	alreadyExistsLaptop := &laptop_message.Laptop{
		Id: savedLaptop.Id,
	}

	testCases := []TestNewLaptopServerTestCase{
		{
			name:   "OK with Id",
			laptop: newLaptopWithId,
			store:  stores.NewInMemoryLaptopStore(),
			code:   codes.OK,
		},
		{
			name:   "OK without Id",
			laptop: newLaptopWithoutId,
			store:  stores.NewInMemoryLaptopStore(),
			code:   codes.OK,
		},
		{
			name:   "InvalidArgument invalid Id",
			laptop: newLaptopWithInvalidId,
			store:  stores.NewInMemoryLaptopStore(),
			code:   codes.InvalidArgument,
		},
		{
			name:   "AlreadyExists",
			laptop: alreadyExistsLaptop,
			store:  savedStore,
			code:   codes.AlreadyExists,
		},
	}

	for i := range testCases {
		testCase := testCases[i]
		test.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			request := &laptop_service.CreateLaptopRequest{
				Laptop: testCase.laptop,
			}

			server := servers.NewLaptopServer(testCase.store)
			response, err := server.CreateLaptop(context.Background(), request)
			if testCase.code == codes.OK {
				require.NoError(t, err)
				require.NotNil(t, response)
				require.NotEmpty(t, response.Id)
			} else {
				require.Error(t, err)
				require.Nil(t, response)
				stt, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, testCase.code, stt.Code())
			}
		})
	}
}
