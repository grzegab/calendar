package http

import (
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

}
