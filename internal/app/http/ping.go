package http

import (
	"net/http"
)

type PingHandler struct{}

func (h *PingHandler) Handle(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong"))
}
