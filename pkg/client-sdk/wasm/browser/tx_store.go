//go:build js && wasm
// +build js,wasm

package browser

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"sync"
	"syscall/js"
	"time"

	"github.com/ark-network/ark/pkg/client-sdk/types"
)

const txsKey = "txs"

type txData struct {
	BoardingTxid string `json:"boarding_txid"`
	RoundTxid    string `json:"round_txid"`
	RedeemTxid   string `json:"redeem_txid"`
	Amount       string `json:"amount"`
	Type         string `json:"type"`
	Settled      string `json:"settled"`
	CreatedAt    string `json:"created_at"`
}

type txStore struct {
	store js.Value

	lock    *sync.Mutex
	eventCh chan types.TransactionEvent
}

func NewTxStore(store js.Value) types.TransactionStore {
	lock := &sync.Mutex{}
	eventCh := make(chan types.TransactionEvent, 10)
	return &txStore{store, lock, eventCh}
}

func (s *txStore) AddTransactions(_ context.Context, txs []types.Transaction) error {
	fmt.Println("ADD TRANSACTIONS")
	if err := s.writeTxs(txs); err != nil {
		return err
	}

	go func(txs []types.Transaction) {
		var eventType types.EventType

		for _, tx := range txs {
			if tx.IsOOR() {
				switch tx.Type {
				case types.TxSent:
					eventType = types.OORSent
				case types.TxReceived:
					eventType = types.OORReceived
				}
				if tx.IsBoarding() {
					eventType = types.BoardingPending
				}
				s.sendEvent(types.TransactionEvent{
					Tx:    tx,
					Event: eventType,
				})
			}
		}
	}(txs)

	return nil
}

func (s *txStore) UpdateTransactions(_ context.Context, txs []types.Transaction) error {
	fmt.Println("UPDATE TRANSACTIONS")
	if err := s.updateTxs(txs); err != nil {
		return err
	}

	go func(txs []types.Transaction) {
		for _, tx := range txs {
			var event types.EventType

			if tx.IsOOR() {
				event = types.OORSettled
			}

			if tx.IsBoarding() {
				event = types.BoardingSettled
			}

			s.sendEvent(types.TransactionEvent{
				Tx:    tx,
				Event: event,
			})
		}
	}(txs)

	return nil
}

func (s *txStore) GetAllTransactions(_ context.Context) ([]types.Transaction, error) {
	rawTxs, err := s.readRawTxs()
	if err != nil {
		return nil, err
	}
	txs := make([]types.Transaction, 0, len(rawTxs))
	for _, rawTx := range rawTxs {
		// nolint:all
		amount, _ := strconv.Atoi(rawTx.Amount)
		settled, _ := strconv.ParseBool(rawTx.Settled)
		createdAt, _ := strconv.Atoi(rawTx.CreatedAt)
		txs = append(txs, types.Transaction{
			TransactionKey: types.TransactionKey{
				BoardingTxid: rawTx.BoardingTxid,
				RoundTxid:    rawTx.RoundTxid,
				RedeemTxid:   rawTx.RedeemTxid,
			},
			Amount:    uint64(amount),
			Type:      types.TxType(rawTx.Type),
			Settled:   settled,
			CreatedAt: time.Unix(int64(createdAt), 0),
		})
	}
	return txs, nil
}

func (s *txStore) GetEventChannel() chan types.TransactionEvent {
	fmt.Println("GET EVENT CH")
	return s.eventCh
}

func (s *txStore) Close() {
	fmt.Println("CLOSE")
	close(s.eventCh)
}

func (s *txStore) sendEvent(event types.TransactionEvent) {
	s.lock.Lock()
	defer s.lock.Unlock()

	select {
	case s.eventCh <- event:
		return
	default:
		fmt.Println("Error: channel is full")
		return
	}
}

func (s *txStore) writeTxs(txs []types.Transaction) error {
	allTxs, err := s.readRawTxs()
	if err != nil {
		return err
	}

	for _, tx := range txs {
		if _, ok := allTxs[tx.TransactionKey.String()]; ok {
			return fmt.Errorf("tx %s already exists", tx.TransactionKey.String())
		}
		rawTx := txData{
			BoardingTxid: tx.BoardingTxid,
			RoundTxid:    tx.RoundTxid,
			RedeemTxid:   tx.RedeemTxid,
			Amount:       strconv.Itoa(int(tx.Amount)),
			Type:         string(tx.Type),
			Settled:      strconv.FormatBool(tx.Settled),
			CreatedAt:    strconv.Itoa(int(tx.CreatedAt.Unix())),
		}
		allTxs[tx.TransactionKey.String()] = rawTx
	}

	buf, err := json.Marshal(allTxs)
	if err != nil {
		return err
	}

	s.store.Call("setItem", txsKey, (string(buf)))
	return nil
}

func (s *txStore) updateTxs(txs []types.Transaction) error {
	allTxs, err := s.readRawTxs()
	if err != nil {
		return err
	}

	for _, tx := range txs {
		_, ok := allTxs[tx.TransactionKey.String()]
		if !ok {
			return fmt.Errorf("tx %s not found", tx.TransactionKey.String())
		}
		var rawTx txData
		buf, err := json.Marshal(tx)
		if err != nil {
			return err
		}
		if err := json.Unmarshal(buf, &rawTx); err != nil {
			return err
		}
		allTxs[tx.TransactionKey.String()] = rawTx
	}

	s.store.Call("setItem", txsKey, allTxs)
	return nil
}

func (s *txStore) readRawTxs() (map[string]txData, error) {
	rawTxs := make(map[string]txData)
	key := s.store.Call("getItem", "txs")
	if key.IsNull() || key.IsUndefined() {
		return rawTxs, nil
	}
	if err := json.Unmarshal([]byte(key.String()), &rawTxs); err != nil {
		return nil, err
	}
	return rawTxs, nil
}
