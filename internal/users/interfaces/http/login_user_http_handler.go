package http

import (
	"context"
	"encoding/json"
	"fmt"
	"github/grzegab/calendar/internal/users/application/login_user"
	"github/grzegab/calendar/internal/users/infrastructure/jwt_generator"
	"net/http"
	"time"
)

type LoginUserHttpHandler struct {
	svc            *login_user.Handler
	tokenGenerator jwt_generator.JwtGenerator
}

func NewLoginHttpHandler(svc *login_user.Handler) *LoginUserHttpHandler {
	return &LoginUserHttpHandler{
		svc: svc,
	}
}

func (h *LoginUserHttpHandler) Handle(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	ctx, cancel := context.WithTimeout(r.Context(), time.Second*5)
	defer cancel()

	var req login_user.LoginInput
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		// log error in some log
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	command := login_user.Query{
		LoginData: req,
	}

	token, err := h.svc.Handle(ctx, command)
	if err != nil {
		// log error in some log
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	body := fmt.Sprintf("bearer: %s", token)

	w.Write([]byte(body))
}
