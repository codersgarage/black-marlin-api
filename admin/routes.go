package admin

import (
	"net/http"

	"github.com/go-chi/chi"
)

var router = chi.NewRouter()

func adminRoutes() http.Handler {
	hr := NewAdminRoutes()
	h := chi.NewRouter()
	h.Group(func(r chi.Router) {
		r.Post("/", hr.saveAdmin)
		r.Get("/", hr.listAdmin)
		r.Get("/{id}", hr.getAdmin)
		r.Patch("/{id}", hr.updateAdmin)
		r.Delete("/{id}", hr.deleteAdmin)
	})
	return h
}
