package app

import (
	"context"
	"database/sql"
	"github/grzegab/calendar/internal/booking"
	bookingPostgres "github/grzegab/calendar/internal/booking/infrastructure/postgres"
	"github/grzegab/calendar/internal/scheduling"
	schedulingPostgres "github/grzegab/calendar/internal/scheduling/infrastructure/postgres"
	"github/grzegab/calendar/internal/shared/infrastructure/auth"
	"github/grzegab/calendar/internal/shared/infrastructure/event_bus"
	"github/grzegab/calendar/internal/shared/interfaces/http/routing"
	"github/grzegab/calendar/internal/users"
	"github/grzegab/calendar/internal/users/application/login_user"
	usersPostgres "github/grzegab/calendar/internal/users/infrastructure"
	"github/grzegab/calendar/internal/users/infrastructure/jwt_generator"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

type App struct {
	db     *sql.DB
	router chi.Router
	// @TODO: ws       *websocket.Server
	jwtVerifier  auth.TokenVerifier
	jwtGenerator jwt_generator.JwtGenerator
	eventBus     event_bus.EventBus

	HealthService *HealthService
	LoginService  *login_user.LoginService

	Users      *users.Module
	Booking    *booking.Module
	Scheduling *scheduling.Module

	registerRoutes []routing.Routable

	server *http.Server

	ReadTimeout  int
	WriteTimeout int
	IdleTimeout  int
}

func CreateApp(opts ...Option) *App {
	app := new(App)
	for _, opt := range opts {
		opt(app)
	}

	app.useEventBus()
	app.loginStrategies()
	app.healthService()
	app.validate()
	app.buildModules()
	app.registerRouting()

	return app
}

func (app *App) Start(addr string) error {
	app.server = &http.Server{
		Addr:         addr,
		Handler:      app.router,
		ReadTimeout:  time.Duration(app.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(app.WriteTimeout) * time.Second,
		IdleTimeout:  time.Duration(app.IdleTimeout) * time.Second,
	}

	return app.server.ListenAndServe()
}

func (app *App) Stop(ctx context.Context) error {
	return app.server.Shutdown(ctx)
}

func (app *App) Router() chi.Router {
	return app.router
}

func (app *App) useEventBus() {
	app.eventBus = event_bus.NewInMemory()
}

func (app *App) loginStrategies() {
	app.LoginService = login_user.NewLoginService(
		&login_user.EmailLoginStrategy{},
		&login_user.PhoneLoginStrategy{},
		&login_user.PasswordLoginStrategy{},
	)
}

func (app *App) healthService() {
	app.HealthService = NewHealthService()
}

func (app *App) buildModules() {
	postgresRepo := usersPostgres.NewUsersRepository(app.db)
	schedulingRepo := schedulingPostgres.NewSchedulingRepository(app.db)
	domainBookingRepo := bookingPostgres.NewBookingRepository(app.db)
	listBookingsRepo := bookingPostgres.NewListBookingsRepository(app.db)

	usersModule := users.NewModule(users.Dependencies{
		UserRepository: postgresRepo,
		LoginService:   app.LoginService,
		TokenVerifier:  app.jwtVerifier,
		TokenGenerator: app.jwtGenerator,
	})

	schedulingModule := scheduling.NewModule(scheduling.Dependencies{
		SchedulingRepository: schedulingRepo,
		TokenVerifier:        app.jwtVerifier,
	})

	bookingModule := booking.NewModule(booking.Dependencies{
		DomainBookingRepository: domainBookingRepo,
		ListBookingsRepository:  listBookingsRepo,
		TokenVerifier:           app.jwtVerifier,
	})

	app.Scheduling = schedulingModule
	app.Users = usersModule
	app.Booking = bookingModule
}

func (app *App) validate() {
	if app.db == nil {
		panic("database is required")
	}

	if app.router == nil {
		panic("router is required")
	}
}

func (app *App) registerRouting() {
	app.HealthService.RegisterRoutes(app.router)
	app.Users.RegisterRoutes(app.router)
	app.Booking.RegisterRoutes(app.router)
	app.Scheduling.RegisterRoutes(app.router)
}
