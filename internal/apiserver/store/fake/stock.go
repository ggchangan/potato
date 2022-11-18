package fake

import (
	"context"
	"github.com/ggchangan/potato/internal/apiserver/model"
	"github.com/ggchangan/potato/internal/pkg/code"
	"github.com/marmotedu/errors"
)

type stocks struct {
	ds *datastore
}

func (s stocks) Get(ctx context.Context, id uint64) (*model.Stock, error) {
	s.ds.RLock()
	defer s.ds.RUnlock()

	for _, st := range s.ds.stocks {
		if st.ID == id {
			return st, nil
		}
	}

	return nil, errors.WithCode(code.ErrStockNotFound, "record not found")
}

func newStocks(ds *datastore) *stocks {
	return &stocks{ds: ds}
}
