package handler

import (
	"Pong/internal/service"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Handler struct {
	service *service.Service
	server  *mux.Router
}

func NewHandler(service *service.Service, server *mux.Router) *Handler {
	return &Handler{service: service, server: server}
}

func (h *Handler) Init() {
	h.server.HandleFunc("/count", h.PongCount).Methods("GET")
}

func (h *Handler) PongCount(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("%d", h.service.Ind)))
}
