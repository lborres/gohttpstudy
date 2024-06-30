package product

import (
	"fmt"
	"net/http"

	m "github.com/lborres/gohttpstudy/internal/middleware"
	"github.com/lborres/gohttpstudy/internal/util"
	"github.com/lborres/gohttpstudy/types"
)

type Handler struct {
	storage types.ProductStorage
}

func NewHandler(storage types.ProductStorage) *Handler {
	return &Handler{storage: storage}
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.Handle("GET /products", m.Logger(h.handleGetAllProducts))
	mux.Handle("GET /products/{id}", m.Logger(h.handleGetProductsByID))
	mux.Handle("POST /products", m.Logger(h.handleAddNewProduct))
	mux.Handle("PUT /products/{id}", m.Logger(h.handleUpdtProduct))
	mux.Handle("DELETE /products/{id}", m.Logger(h.handleDelProduct))
}

func (h *Handler) handleGetAllProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.storage.GetAllProducts()
	if err != nil {
		util.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	util.EncodeJSON(w, http.StatusOK, products)
}

func (h *Handler) handleGetProductsByID(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Placeholder")
}

// Admin only
func (h *Handler) handleAddNewProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Placeholder")
}

func (h *Handler) handleUpdtProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Placeholder")
}

func (h *Handler) handleDelProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Placeholder")
}
