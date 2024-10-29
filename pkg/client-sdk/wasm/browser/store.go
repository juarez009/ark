//go:build js && wasm
// +build js,wasm

package browser

import (
	"syscall/js"

	"github.com/ark-network/ark/pkg/client-sdk/types"
)

// TODO: support vtxo and transaction stores localstorage impls.
type localStorageStore struct {
	configStore types.ConfigStore
	txStore     types.TransactionStore
}

func NewLocalStorageStore() types.Store {
	store := js.Global().Get("localStorage")
	configStore := NewConfigStore(store)
	txStore := NewTxStore(store)
	return &localStorageStore{configStore, txStore}
}

func (s *localStorageStore) ConfigStore() types.ConfigStore {
	return s.configStore
}

func (s *localStorageStore) VtxoStore() types.VtxoStore {
	return nil
}

func (s *localStorageStore) TransactionStore() types.TransactionStore {
	return s.txStore
}

func (s *localStorageStore) Close() {}
