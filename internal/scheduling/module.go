package scheduling

import (
	"github/grzegab/calendar/internal/scheduling/application/confirm_schedule"
	"github/grzegab/calendar/internal/scheduling/application/decline_schedule"
	"github/grzegab/calendar/internal/scheduling/application/new_timeslot"
	"github/grzegab/calendar/internal/scheduling/application/schedule_details"
	"github/grzegab/calendar/internal/scheduling/application/schedule_list"
	"github/grzegab/calendar/internal/scheduling/domain"
	"github/grzegab/calendar/internal/scheduling/interfaces/http"
	"github/grzegab/calendar/internal/shared/infrastructure/auth"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Dependencies struct {
	SchedulingRepository domain.SchedulingRepository
	TokenVerifier        auth.TokenVerifier
}

type Module struct {
	listSlotSvc                *schedule_list.Handler
	listTimeslotsHttpHandler   *http.ListSlotsHttpHandler
	timeslotDetailsSvc         *schedule_details.Handler
	timeslotDetailsHttpHandler *http.SlotDetailsHttpHandler
	newTimeslotService         *new_timeslot.Handler
	newTimeslotHttpHandler     *http.NewSlotHttpHandler
	confirmScheduleService     *confirm_schedule.Handler
	confirmScheduleHttpHandler *http.ConfirmScheduleHttpHandler
	declineScheduleService     *decline_schedule.Handler
	declineScheduleHttpHandler *http.DeclineScheduleHttpHandler
	tokenVerifier              auth.TokenVerifier
}

func NewModule(dep Dependencies) *Module {
	// query
	listSchedulesSvc := schedule_list.NewHandler(dep.SchedulingRepository)
	listTimeslotsHttpHandler := http.NewListSlotsHttpHandler(listSchedulesSvc)
	timeslotDetailsSvc := schedule_details.NewHandler(dep.SchedulingRepository)
	timeslotDetailsHttpHandler := http.NewSlotDetailsHttpHandler(timeslotDetailsSvc)

	// commands
	newTimeslotSvc := new_timeslot.NewHandler(dep.SchedulingRepository)
	newTimeslotHttpHandler := http.CreateNewSlotHttpHandler(newTimeslotSvc)
	confirmScheduleSvc := confirm_schedule.NewHandler(dep.SchedulingRepository)
	confirmScheduleHttpHandler := http.CreateConfirmScheduleHttpHandler(confirmScheduleSvc)
	declineScheduleSvc := decline_schedule.NewHandler(dep.SchedulingRepository)
	declineScheduleHttpHandler := http.CreateDeclineScheduleHttpHandler(declineScheduleSvc)

	return &Module{
		listSlotSvc:                listSchedulesSvc,
		listTimeslotsHttpHandler:   listTimeslotsHttpHandler,
		timeslotDetailsSvc:         timeslotDetailsSvc,
		timeslotDetailsHttpHandler: timeslotDetailsHttpHandler,
		newTimeslotService:         newTimeslotSvc,
		newTimeslotHttpHandler:     newTimeslotHttpHandler,
		confirmScheduleService:     confirmScheduleSvc,
		confirmScheduleHttpHandler: confirmScheduleHttpHandler,
		declineScheduleService:     declineScheduleSvc,
		declineScheduleHttpHandler: declineScheduleHttpHandler,
		tokenVerifier:              dep.TokenVerifier,
	}
}

func (m *Module) RegisterRoutes(r chi.Router) {
	r.Route("/scheduling", func(r chi.Router) {
		r.Use(auth.JwtMiddleware(m.tokenVerifier))
		r.Use(middleware.Recoverer)

		r.Get("/", m.listTimeslotsHttpHandler.Handle)
		r.Get("/{id}", m.timeslotDetailsHttpHandler.Handle)
		r.Post("/", m.newTimeslotHttpHandler.Handle)
		r.Post("/{id}", m.confirmScheduleHttpHandler.Handle)
		r.Delete("/{id}", m.declineScheduleHttpHandler.Handle)
	})
}
