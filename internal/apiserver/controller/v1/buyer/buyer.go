package buyer

import (
	"github.com/ggchangan/potato/internal/apiserver/service"
	"github.com/ggchangan/potato/internal/apiserver/store"
)

// BuyerController create a stock handler used to handle request for stock resource.
type BuyerController struct {
	srv service.Service
}

// NewBuyerController creates a stock handler.
func NewBuyerController(store store.Factory) *BuyerController {
	return &BuyerController{
		srv: service.NewService(store),
	}
}
