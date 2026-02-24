package http

import (
	"context"
	"encoding/json"
	"github/grzegab/calendar/internal/scheduling/application/new_timeslot"
	"github/grzegab/calendar/internal/shared/infrastructure/auth"
	"net/http"
	"time"
)

type NewSlotHttpHandler struct {
	service *new_timeslot.Handler
}

func CreateNewSlotHttpHandler(svc *new_timeslot.Handler) *NewSlotHttpHandler {
	return &NewSlotHttpHandler{service: svc}
}

// Handle create new timeslot
func (h *NewSlotHttpHandler) Handle(w http.ResponseWriter, r *http.Request) {
	teacherID, ok := auth.UserIDFromContext(r.Context())
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	// get data from request
	var req struct {
		StartTime time.Time `json:"start_time"`
		EndTime   time.Time `json:"end_time"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.StartTime.IsZero() || req.EndTime.IsZero() {
		http.Error(w, "Invalid start or end time", http.StatusBadRequest)
		return
	}

	if teacherID == "" {
		http.Error(w, "Invalid teacher ID", http.StatusBadRequest)
		return
	}

	command := new_timeslot.Command{
		TeacherID: teacherID,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
	}

	ctx := context.WithValue(r.Context(), "teacher_id", teacherID)

	err := h.service.Handle(ctx, command)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
