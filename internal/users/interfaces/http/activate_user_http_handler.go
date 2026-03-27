package http

import (
	"github/grzegab/calendar/internal/shared/infrastructure/auth"
	"github/grzegab/calendar/internal/users/application/activate_user"
	"net/http"
)

type ActivateUserHttpHandler struct {
	svc *activate_user.Handler
}

func NewActivateHttpHandler(svc *activate_user.Handler) *ActivateUserHttpHandler {
	return &ActivateUserHttpHandler{svc: svc}
}

func (h *ActivateUserHttpHandler) Handle(w http.ResponseWriter, r *http.Request) {
	userID, ok := auth.UserIDFromContext(r.Context())
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	if userID != "123" { // Only admin can activate user, make roles
		http.Error(w, "unauthorized", http.StatusForbidden)
	}

	http.Error(w, "not initialized", http.StatusNotImplemented)
}
