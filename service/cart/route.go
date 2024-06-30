package cart

import (
	"fmt"
	"net/http"

	"github.com/lborres/gohttpstudy/types"
)

type Handler struct {
	storage types.CartStorage
}

func NewHandler(storage types.CartStorage) *Handler {
	return &Handler{storage: storage}
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /cart", h.handleGetCart)
	mux.HandleFunc("POST /cart", h.handleAddCart)
	mux.HandleFunc("PUT /cart", h.handleUpdateCart)
	mux.HandleFunc("DELETE /cart/{id}", h.handleDelCart)
}

func (h *Handler) handleGetCart(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Placeholder")
}

func (h *Handler) handleAddCart(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Placeholder")
}

func (h *Handler) handleUpdateCart(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Placeholder")
}

func (h *Handler) handleDelCart(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Placeholder")
}
