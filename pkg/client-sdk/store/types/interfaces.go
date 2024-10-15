package storetypes

import "context"

type Store interface {
	ConfigStore() ConfigStore
	TransactionStore() TransactionStore
	VtxoStore() VtxoStore
	Close()
}

type ConfigStore interface {
	GetType() string
	GetDatadir() string
	AddData(ctx context.Context, data ConfigData) error
	GetData(ctx context.Context) (*ConfigData, error)
	CleanData(ctx context.Context) error
	Close()
}

type TransactionStore interface {
	AddTransactions(ctx context.Context, txs []Transaction) error
	UpdateTransactions(ctx context.Context, txs []Transaction) error
	GetAllTransactions(ctx context.Context) ([]Transaction, error)
	GetEventChannel() chan TransactionEvent
	GetBoardingTxs(ctx context.Context) ([]Transaction, error)
	Close()
}

type VtxoStore interface {
	AddVtxos(ctx context.Context, vtxos []Vtxo) error
	UpdateVtxos(ctx context.Context, vtxos []Vtxo) error
	GetAllVtxos(ctx context.Context) (spendable []Vtxo, spent []Vtxo, err error)
	GetVtxos(ctx context.Context, keys []VtxoKey) ([]Vtxo, error)
	Close()
}
