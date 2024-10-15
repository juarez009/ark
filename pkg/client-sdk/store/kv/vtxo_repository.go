package badgerstore

import (
	"context"
	"errors"
	"fmt"
	"path/filepath"

	storetypes "github.com/ark-network/ark/pkg/client-sdk/store/types"
	"github.com/dgraph-io/badger/v4"
	log "github.com/sirupsen/logrus"
	"github.com/timshannon/badgerhold/v4"
)

const (
	vtxoStoreDir = "vtxos"
)

type vtxoRepository struct {
	db *badgerhold.Store
}

func NewVtxoStore(dir string, logger badger.Logger) (storetypes.VtxoStore, error) {
	badgerDb, err := createDB(filepath.Join(dir, vtxoStoreDir), logger)
	if err != nil {
		return nil, fmt.Errorf("failed to open round events store: %s", err)
	}
	return &vtxoRepository{badgerDb}, nil
}

func (v *vtxoRepository) AddVtxos(ctx context.Context, vtxos []storetypes.Vtxo) error {
	for _, vtxo := range vtxos {
		if err := v.db.Insert(vtxo.VtxoKey.String(), &vtxo); err != nil {
			return err
		}
	}
	return nil
}

func (v *vtxoRepository) UpdateVtxos(ctx context.Context, vtxos []storetypes.Vtxo) error {
	for _, vtxo := range vtxos {
		if err := v.db.Update(vtxo.VtxoKey.String(), &vtxo); err != nil {
			return err
		}
	}
	return nil
}

func (v *vtxoRepository) GetAllVtxos(
	ctx context.Context,
) (spendable []storetypes.Vtxo, spent []storetypes.Vtxo, err error) {
	var allVtxos []storetypes.Vtxo
	err = v.db.Find(&allVtxos, nil)
	if err != nil {
		return nil, nil, err
	}

	for _, vtxo := range allVtxos {
		if vtxo.Spent {
			spent = append(spent, vtxo)
		} else {
			spendable = append(spendable, vtxo)
		}
	}
	return
}

func (v *vtxoRepository) GetVtxos(
	ctx context.Context,
	keys []storetypes.VtxoKey,
) ([]storetypes.Vtxo, error) {
	var vtxos []storetypes.Vtxo
	for _, key := range keys {
		var vtxo storetypes.Vtxo
		err := v.db.Get(key.String(), &vtxo)
		if err != nil {
			if errors.Is(err, badgerhold.ErrNotFound) {
				continue
			}

			return nil, err
		}
		vtxos = append(vtxos, vtxo)
	}

	return vtxos, nil
}

func (v *vtxoRepository) Close() {
	if err := v.db.Close(); err != nil {
		log.Debugf("error on closing db: %s", err)
	}
}
