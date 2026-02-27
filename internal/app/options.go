package app

import (
	"database/sql"
	"github/grzegab/calendar/internal/shared/infrastructure/auth"

	"github.com/go-chi/chi/v5"
)

type Option func(*App)

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

//func WithWebSocket(ws *websocket.Server) Option {
//	return func(app *App) {
//		app.ws = ws
//	}
//}

func WithAuthVerifier(secret string) Option {
	return func(app *App) {
		app.verifier = auth.NewVerifier(auth.HMACKeyFunc([]byte(secret)))
	}
}
