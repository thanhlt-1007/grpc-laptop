package stores

import (
	"errors"
	"fmt"
	"sync"

	"grpc-laptop/go_protos/messages/laptop_message"

	"github.com/jinzhu/copier"
)

type LaptopStore interface {
	Save(laptop *laptop_message.Laptop) error
	Find(id string) (*laptop_message.Laptop, error)
}

var ErrAlreadyExists = errors.New("AlreadyExists")

type InMemoryLaptopStore struct {
	mutex sync.Mutex

	data map[string]*laptop_message.Laptop
}

func NewInMemoryLaptopStore() *InMemoryLaptopStore {
	store := InMemoryLaptopStore{
		data: make(map[string]*laptop_message.Laptop),
	}
	return &store
}

func (store *InMemoryLaptopStore) Save(laptop *laptop_message.Laptop) error {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	if store.data[laptop.Id] != nil {
		return ErrAlreadyExists
	}

	newLaptop := &laptop_message.Laptop{}
	err := copier.Copy(newLaptop, laptop)
	if err != nil {
		return fmt.Errorf("Can't copy laptop %#v", err)
	}

	store.data[newLaptop.Id] = newLaptop
	return nil
}

func (store *InMemoryLaptopStore) Find(id string) (*laptop_message.Laptop, error) {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	laptop := store.data[id]
	if laptop == nil {
		return nil, nil
	}

	copyLaptop := &laptop_message.Laptop{}
	err := copier.Copy(copyLaptop, laptop)
	if err != nil {
		return nil, fmt.Errorf("Can't copy laptop %#v", err)
	}

	return copyLaptop, nil
}
