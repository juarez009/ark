package store

import (
	"fmt"

	filestore "github.com/ark-network/ark/pkg/client-sdk/store/file"
	inmemorystore "github.com/ark-network/ark/pkg/client-sdk/store/inmemory"
	kvstore "github.com/ark-network/ark/pkg/client-sdk/store/kv"
	storetypes "github.com/ark-network/ark/pkg/client-sdk/store/types"
	log "github.com/sirupsen/logrus"
)

const (
	InMemoryStore = "inmemory"
	FileStore     = "file"
	KVStore       = "kv"
)

type service struct {
	configStore storetypes.ConfigStore
	vtxoStore   storetypes.VtxoStore
	txStore     storetypes.TransactionStore
}

type Config struct {
	ConfigStoreType  string
	AppDataStoreType string

	BaseDir string
}

func NewStore(storeConfig Config) (storetypes.Store, error) {
	var (
		configStore storetypes.ConfigStore
		vtxoStore   storetypes.VtxoStore
		txStore     storetypes.TransactionStore
		err         error

		dir = storeConfig.BaseDir
	)

	switch storeConfig.ConfigStoreType {
	case InMemoryStore:
		configStore, err = inmemorystore.NewConfigStore()
	case FileStore:
		configStore, err = filestore.NewConfigStore(dir)
	default:
		err = fmt.Errorf("unknown config store type")
	}
	if err != nil {
		return nil, err
	}

	switch storeConfig.AppDataStoreType {
	case KVStore:
		logger := log.New()
		vtxoStore, err = kvstore.NewVtxoStore(dir, logger)
		if err != nil {
			return nil, err
		}
		txStore, err = kvstore.NewTransactionStore(dir, logger)
	default:
		err = fmt.Errorf("unknown app data store type")
	}
	if err != nil {
		return nil, err
	}

	return &service{configStore, vtxoStore, txStore}, nil
}

func (s *service) ConfigStore() storetypes.ConfigStore {
	return s.configStore
}

func (s *service) VtxoStore() storetypes.VtxoStore {
	return s.vtxoStore
}

func (s *service) TransactionStore() storetypes.TransactionStore {
	return s.txStore
}

func (s *service) Close() {
	s.configStore.Close()
	s.vtxoStore.Close()
	s.txStore.Close()
}
