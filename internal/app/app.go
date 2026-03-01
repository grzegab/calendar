package app

import (
	"database/sql"
	"github/grzegab/calendar/internal/booking"
	bookingPostgres "github/grzegab/calendar/internal/booking/infrastructure/postgres"
	"github/grzegab/calendar/internal/scheduling"
	schedulingPostgres "github/grzegab/calendar/internal/scheduling/infrastructure/postgres"
	"github/grzegab/calendar/internal/shared/infrastructure/auth"
	"github/grzegab/calendar/internal/shared/interfaces/http/routing"
	"github/grzegab/calendar/internal/users"
	usersPostgres "github/grzegab/calendar/internal/users/infrastructure/postgres"

	"github.com/go-chi/chi/v5"
)

type App struct {
	db     *sql.DB
	router chi.Router
	// @TODO: ws       *websocket.Server
	verifier auth.TokenVerifier

	Health     *HealthService
	Users      *users.Module
	Booking    *booking.Module
	Scheduling *scheduling.Module

	registerRoutes []routing.Routable
}

func CreateApp(opts ...Option) *App {
	app := new(App)
	for _, opt := range opts {
		opt(app)
	}

	app.validate()
	app.buildModules()
	app.registerRouting()

	return &App{}
}

func (a *App) Router() chi.Router {
	return a.router
}

func (a *App) buildModules() {
	postgresRepo := usersPostgres.NewUsersRepository(a.db)
	schedulingRepo := schedulingPostgres.NewSchedulingRepository(a.db)
	domainBookingRepo := bookingPostgres.NewBookingRepository(a.db)
	listBookingsRepo := bookingPostgres.NewListBookingsRepository(a.db)

	usersModule := users.NewModule(users.Dependencies{
		UserRepository: postgresRepo,
	})

	schedulingModule := scheduling.NewModule(scheduling.Dependencies{
		SchedulingRepository: schedulingRepo,
		TokenVerifier:        a.verifier,
	})

	bookingModule := booking.NewModule(booking.Dependencies{
		DomainBookingRepository: domainBookingRepo,
		ListBookingsRepository:  listBookingsRepo,
	})

	a.Health = NewHealthService()
	a.Scheduling = schedulingModule
	a.Users = usersModule
	a.Booking = bookingModule
}

func (a *App) validate() {
	if a.db == nil {
		panic("database is required")
	}

	if a.router == nil {
		panic("router is required")
	}
}

func (a *App) registerRouting() {
	a.Health.RegisterRoutes(a.router)
	a.Users.RegisterRoutes(a.router)
	a.Booking.RegisterRoutes(a.router)
}
