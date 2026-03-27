package users

import (
	"github/grzegab/calendar/internal/shared/infrastructure/auth"
	"github/grzegab/calendar/internal/users/application/activate_user"
	"github/grzegab/calendar/internal/users/application/login_user"
	"github/grzegab/calendar/internal/users/application/register_user"
	"github/grzegab/calendar/internal/users/application/registered_user_list"
	"github/grzegab/calendar/internal/users/application/unregistered_user_list"
	"github/grzegab/calendar/internal/users/domain"
	"github/grzegab/calendar/internal/users/infrastructure/jwt_generator"
	"github/grzegab/calendar/internal/users/interfaces/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Module struct {
	activateUserHttpHandler         *http.ActivateUserHttpHandler
	activateUserService             *activate_user.Handler
	registeredUserListHttpHandler   *http.RegisteredUsersListHttpHandler
	registeredUserListService       *registered_user_list.Handler
	unregisteredUserListHttpHandler *http.UnregisteredUsersListHttpHandler
	unregisteredUserListService     *unregistered_user_list.Handler
	loginUserHttpHandler            *http.LoginUserHttpHandler
	loginUserService                *login_user.Handler
	registerUserHttpHandler         *http.RegisterUserHttpHandler
	registerUserService             *register_user.Handler
	tokenVerifier                   auth.TokenVerifier
	tokenGenerator                  jwt_generator.JwtGenerator
}

type Dependencies struct {
	RegisteredUserRepo   registered_user_list.ReadRepository
	UnregisteredUserRepo unregistered_user_list.ReadRepository
	UserRepository       domain.Repository
	LoginService         *login_user.LoginService
	TokenVerifier        auth.TokenVerifier
	TokenGenerator       jwt_generator.JwtGenerator
}

func NewModule(dep Dependencies) *Module {
	// query
	registeredUserListService := registered_user_list.NewHandler(dep.RegisteredUserRepo)
	registeredUserListHttpHandler := http.NewRegisteredUsersListHttpHandler(registeredUserListService)
	unregisteredUserListService := unregistered_user_list.NewHandler(dep.UnregisteredUserRepo)
	unregisteredUserListHttpHandler := http.NewUnregisteredUsersListHttpHandler(unregisteredUserListService)

	// command
	registerUserService := register_user.NewHandler(dep.UserRepository)
	registerUserHttpHandler := http.NewRegisterUserHttpHandler(registerUserService)
	loginUserService := login_user.NewHandler(dep.UserRepository, dep.LoginService, dep.TokenGenerator)
	loginUserHttpHandler := http.NewLoginHttpHandler(loginUserService)
	activateUserService := activate_user.NewHandler(dep.UserRepository)
	activateUserHttpHandler := http.NewActivateHttpHandler(activateUserService)

	return &Module{
		registeredUserListService:       registeredUserListService,
		registeredUserListHttpHandler:   registeredUserListHttpHandler,
		unregisteredUserListService:     unregisteredUserListService,
		unregisteredUserListHttpHandler: unregisteredUserListHttpHandler,
		registerUserService:             registerUserService,
		registerUserHttpHandler:         registerUserHttpHandler,
		loginUserService:                loginUserService,
		loginUserHttpHandler:            loginUserHttpHandler,
		activateUserService:             activateUserService,
		activateUserHttpHandler:         activateUserHttpHandler,
	}
}

func (m *Module) RegisterRoutes(r chi.Router) {
	r.Route("/users", func(r chi.Router) {
		r.Use(middleware.Recoverer)

		r.Post("/register", m.registerUserHttpHandler.Handle)
		r.Post("/login", m.loginUserHttpHandler.Handle)

		r.Group(func(r chi.Router) {
			r.Use(auth.JwtMiddleware(m.tokenVerifier))

			r.Get("/registered", m.registeredUserListHttpHandler.Handle)
			r.Get("/unregistered", m.unregisteredUserListHttpHandler.Handle)
			r.Patch("/{id}", m.activateUserHttpHandler.Handle)
		})
	})
}
