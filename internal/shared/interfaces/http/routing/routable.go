package routing

import "github.com/go-chi/chi/v5"

type Routable interface {
	RegisterRoutes(r chi.Router)
}
