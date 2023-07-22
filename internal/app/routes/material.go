package routes

import (
	"github.com/azka1415/crud-bootcamp-evermos/internal/app/handlers"
	"github.com/go-chi/chi"
)

func SetMaterialRoutes(r chi.Router) {
	r.Route("/materials", func(r chi.Router) {
		r.Get("/", handlers.HandleGetMaterial)
		r.Post("/", handlers.HandlePostMaterial)
		r.Get("/{id}", handlers.HandleGetMaterialByID)
		r.Put("/{id}", handlers.HandleUpdateMaterial)
		r.Delete("/{id}", handlers.HandleDeleteMaterial)
	})
}
