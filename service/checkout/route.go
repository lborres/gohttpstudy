package checkout

import (
	"net/http"

	"github.com/lborres/gohttpstudy/types"
)

type Handler struct {
	storage types.CheckoutStorage
}

func NewHandler(storage types.CartStorage) *Handler {
	return &Handler{storage: storage}
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {

}
