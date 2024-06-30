package misc

import (
	"fmt"
	"net/http"

	m "github.com/lborres/gohttpstudy/internal/middleware"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.Handle("GET /ping", m.Logger(h.handlePing))
	mux.Handle("GET /about", m.Logger(h.handleAbout))
}

func (h *Handler) handlePing(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "pong")
}

func (h *Handler) handleAbout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Simulated e-commerce site - Go HTTP Server Study by Louis Borres")
}
