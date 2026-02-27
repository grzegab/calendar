package users

import (
	"github/grzegab/calendar/internal/users/application/activate_user"
	"github/grzegab/calendar/internal/users/application/active_user_list"
	"github/grzegab/calendar/internal/users/application/login_user"
	"github/grzegab/calendar/internal/users/application/register_user"
	"github/grzegab/calendar/internal/users/application/user_details"
	"github/grzegab/calendar/internal/users/domain"
	"github/grzegab/calendar/internal/users/interfaces/http"

	"github.com/go-chi/chi/v5"
)

type Module struct {
	activateUserHttpHandler   *http.ActivateUserHttpHandler
	activateUserService       *activate_user.Handler
	activeUserListHttpHandler *http.ActiveUsersHttpHandler
	activeUserListService     *active_user_list.Handler
	loginUserHttpHandler      *http.LoginUserHttpHandler
	loginUserService          *login_user.Handler
	registerUserHttpHandler   *http.RegisterUserHttpHandler
	registerUserService       *register_user.Handler
	userDetailsHttpHandler    *http.UserDetailsHttpHandler
	userDetailsService        *user_details.Handler
}

type Dependencies struct {
	ActiveUserRepository  active_user_list.ReadRepository
	UserDetailsRepository user_details.ReadRepository
	UserRepository        domain.Repository
}

func NewModule(dep Dependencies) *Module {
	// query
	activeUserListService := active_user_list.NewHandler(dep.ActiveUserRepository)
	activeUserListHttpHandler := http.NewActiveUsersHandler(activeUserListService)
	userDetailsService := user_details.NewHandler(dep.UserDetailsRepository)
	userDetailsHttpHandler := http.NewUserDetailsHttpHandler(userDetailsService)

	// command
	registerUserService := register_user.NewHandler(dep.UserRepository)
	registerUserHttpHandler := http.NewRegisterUserHttpHandler(registerUserService)
	loginUserService := login_user.NewHandler(dep.UserRepository)
	loginUserHttpHandler := http.NewLoginHttpHandler(loginUserService)
	activateUserService := activate_user.NewHandler(dep.UserRepository)
	activateUserHttpHandler := http.NewActivateHttpHandler(activateUserService)

	return &Module{
		activeUserListService:     activeUserListService,
		activeUserListHttpHandler: activeUserListHttpHandler,
		userDetailsService:        userDetailsService,
		userDetailsHttpHandler:    userDetailsHttpHandler,
		registerUserService:       registerUserService,
		registerUserHttpHandler:   registerUserHttpHandler,
		loginUserService:          loginUserService,
		loginUserHttpHandler:      loginUserHttpHandler,
		activateUserService:       activateUserService,
		activateUserHttpHandler:   activateUserHttpHandler,
	}
}

func (m *Module) RegisterRoutes(r chi.Router) {
	r.Route("/users", func(r chi.Router) {
		r.Get("/", m.activeUserListHttpHandler.Handle)
		r.Get("/{id}", m.activeUserListHttpHandler.Handle)

		r.Post("/register", m.registerUserHttpHandler.Handle)
		r.Post("/login", m.loginUserHttpHandler.Handle)
		r.Patch("/{id}", m.activateUserHttpHandler.Handle)
	})
}
