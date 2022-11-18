package stock

import (
	"github.com/ggchangan/potato/internal/apiserver/service"
	"github.com/ggchangan/potato/internal/apiserver/store"
)

// StockController create a stock handler used to handle request for stock resource.
type StockController struct {
	srv service.Service
}

// NewStockController creates a stock handler.
func NewStockController(store store.Factory) *StockController {
	return &StockController{
		srv: service.NewService(store),
	}
}
