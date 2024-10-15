package badgerstore

import (
	"context"
	"fmt"
	"path/filepath"
	"sort"

	storetypes "github.com/ark-network/ark/pkg/client-sdk/store/types"
	"github.com/dgraph-io/badger/v4"
	log "github.com/sirupsen/logrus"
	"github.com/timshannon/badgerhold/v4"
)

const (
	transactionStoreDir = "transactions"
)

type transactionRepository struct {
	db      *badgerhold.Store
	eventCh chan storetypes.TransactionEvent
}

func NewTransactionStore(
	dir string, logger badger.Logger,
) (storetypes.TransactionStore, error) {
	badgerDb, err := createDB(filepath.Join(dir, transactionStoreDir), logger)
	if err != nil {
		return nil, fmt.Errorf("failed to open round events store: %s", err)
	}
	return &transactionRepository{
		db:      badgerDb,
		eventCh: make(chan storetypes.TransactionEvent),
	}, nil
}

func (t *transactionRepository) GetBoardingTxs(ctx context.Context) ([]storetypes.Transaction, error) {
	var txs []storetypes.Transaction
	query := badgerhold.Where("BoardingTxid").Ne("")
	err := t.db.Find(&txs, query)
	return txs, err
}

func (t *transactionRepository) AddTransactions(
	ctx context.Context, txs []storetypes.Transaction,
) error {
	for _, tx := range txs {
		if err := t.db.Insert(tx.TransactionKey.String(), &tx); err != nil {
			return err
		}
		go func(tx storetypes.Transaction) {
			var eventType storetypes.EventType

			if tx.IsOOR() {
				switch tx.Type {
				case storetypes.TxSent:
					eventType = storetypes.OORSent
				case storetypes.TxReceived:
					eventType = storetypes.OORReceived
				}
			}

			if tx.IsBoarding() {
				eventType = storetypes.BoardingPending
			}

			t.eventCh <- storetypes.TransactionEvent{
				Tx:    tx,
				Event: eventType,
			}
		}(tx)
	}
	return nil
}

func (t *transactionRepository) UpdateTransactions(
	ctx context.Context, txs []storetypes.Transaction,
) error {
	for _, tx := range txs {
		if err := t.db.Upsert(tx.TransactionKey.String(), &tx); err != nil {
			return err
		}
		go func(tx storetypes.Transaction) {
			var event storetypes.EventType

			if tx.IsOOR() {
				event = storetypes.OORSettled
			}

			if tx.IsBoarding() {
				event = storetypes.BoardingSettled
			}

			t.eventCh <- storetypes.TransactionEvent{
				Tx:    tx,
				Event: event,
			}
		}(tx)
	}
	return nil
}

func (t *transactionRepository) GetAllTransactions(
	ctx context.Context,
) ([]storetypes.Transaction, error) {
	var txs []storetypes.Transaction
	err := t.db.Find(&txs, nil)

	sort.Slice(txs, func(i, j int) bool {
		txi := txs[i]
		txj := txs[j]
		if txi.CreatedAt.Equal(txj.CreatedAt) {
			return txi.Type > txj.Type
		}
		return txi.CreatedAt.After(txj.CreatedAt)
	})

	return txs, err
}

func (t *transactionRepository) GetEventChannel() chan storetypes.TransactionEvent {
	return t.eventCh
}

func (t *transactionRepository) Close() {
	if err := t.db.Close(); err != nil {
		log.Debugf("error on closing transactions db: %s", err)
	}
	close(t.eventCh)
}
