package inmemorystore

import (
	"context"
	"sync"

	storetypes "github.com/ark-network/ark/pkg/client-sdk/store/types"
)

type store struct {
	data *storetypes.ConfigData
	lock *sync.RWMutex
}

func NewConfigStore() (storetypes.ConfigStore, error) {
	lock := &sync.RWMutex{}
	return &store{lock: lock}, nil
}

func (s *store) Close() {}

func (s *store) GetType() string {
	return "inmemory"
}

func (s *store) GetDatadir() string {
	return ""
}

func (s *store) AddData(
	_ context.Context, data storetypes.ConfigData,
) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.data = &data
	return nil
}

func (s *store) GetData(_ context.Context) (*storetypes.ConfigData, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	if s.data == nil {
		return nil, nil
	}

	return s.data, nil
}

func (s *store) CleanData(_ context.Context) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.data = nil
	return nil
}
