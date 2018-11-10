package admin

import (
	"net/http"

	"github.com/go-chi/chi"
)

var router = chi.NewRouter()

// Routes ... endpoint definition
func Routes() http.Handler {
	hr := NewRoutes()
	h := chi.NewRouter()
	h.Group(func(r chi.Router) {
		r.Post("/", hr.save)
		r.Get("/", hr.list)
		r.Get("/{id}", hr.get)
		r.Patch("/{id}", hr.update)
		r.Delete("/{id}", hr.delete)
	})
	return h
}
