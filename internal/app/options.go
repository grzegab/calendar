package app

import (
	"database/sql"
	"github/grzegab/calendar/internal/shared/infrastructure/auth"
	"github/grzegab/calendar/internal/users/infrastructure/jwt_generator"

	"github.com/go-chi/chi/v5"
)

type Option func(*App)

func WithTimeoutConfig(timeConfig HttpConfig) Option {
	return func(app *App) {
		app.ReadTimeout = timeConfig.ReadTimeout
		app.WriteTimeout = timeConfig.WriteTimeout
		app.IdleTimeout = timeConfig.IdleTimeout
	}
}

func WithDB(db *sql.DB) Option {
	return func(app *App) {
		app.db = db
	}
}

func WithRouter(router chi.Router) Option {
	return func(app *App) {
		app.router = router
	}
}

func WithJwtGenerator(secret string) Option {
	return func(app *App) {
		app.jwtGenerator = jwt_generator.NewJwtGenerator(secret)
	}
}

//func WithWebSocket(ws *websocket.Server) Option {
//	return func(app *App) {
//		app.ws = ws
//	}
//}

func WithAuthVerifier(secret string) Option {
	return func(app *App) {
		app.jwtVerifier = auth.NewVerifier(auth.HMACKeyFunc([]byte(secret)))
	}
}
