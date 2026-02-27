package router

import (
	"net/http"

	"github/grzegab/calendar/internal/shared/ws"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type Router struct {
	chi *chi.Mux
	Hub *ws.Hub
}

// Handler returns the underlying chi router as an http.Handler
func (r *Router) Handler() chi.Router {
	return r.chi
}

// New returns a new router instance witch chi router
func New(origins []string, h *ws.Hub) *Router {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:     origins,
		AllowOriginFunc:    nil,
		AllowedMethods:     nil,
		AllowedHeaders:     nil,
		ExposedHeaders:     nil,
		AllowCredentials:   false,
		MaxAge:             0,
		OptionsPassthrough: false,
		Debug:              false,
	}))

	//// heartbeat
	//r.Get("/ping", to_delete.Ping)
	//
	//// html plain views
	//r.Get("/login", to_delete.Login)
	//r.Post("/login", to_delete.LoginPost)
	//r.Get("/", to_delete.Calendar)
	//r.Get("/settings", to_delete.Settings)
	//r.Post("/settings", to_delete.SettingsPost)
	//
	// websocket paths
	r.Get("/ws", func(w http.ResponseWriter, r *http.Request) {
		ws.WsUpgrade(h, w, r)
	})

	return &Router{chi: r, Hub: h}
}
