package user

import (
	"net/http"

	"github.com/FrancoRutigliano/internal/types"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{
		store: store,
	}
}

func (h *Handler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("GET /login", h.handleLogin)
	router.HandleFunc("POST /register", h.handleRegister)
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Login"))
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Register"))
}
