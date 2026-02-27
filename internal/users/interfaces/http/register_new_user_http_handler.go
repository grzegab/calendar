package http

import (
	"github/grzegab/calendar/internal/users/application/register_user"
	"net/http"
)

type RegisterUserHttpHandler struct {
	svc *register_user.Handler
}

func NewRegisterUserHttpHandler(svc *register_user.Handler) *RegisterUserHttpHandler {
	return &RegisterUserHttpHandler{svc: svc}
}

func (h *RegisterUserHttpHandler) Handle(w http.ResponseWriter, r *http.Request) {
	//var req struct {
	//	Email string `json:"email"`
	//	Phone string `json:"phone"`
	//}
	//
	//if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
	//	http.Error(w, "Invalid request body", http.StatusBadRequest)
	//	return
	//}
	//
	//user, err := h.svc.Create(r.Context(), req.Email, req.Phone)
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}
	//
	//w.WriteHeader(http.StatusCreated)
	//json.NewEncoder(w).Encode(user)
}
