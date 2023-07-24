package routes

import (
	"github.com/azka1415/crud-bootcamp-evermos/internal/app/handlers"
	"github.com/go-chi/chi"
)

func SetMaterialRoutes(r chi.Router) {
	matHandle := handlers.NewMaterialHandler()
	r.Route("/materials", func(r chi.Router) {
		r.Get("/", matHandle.GetMaterial)
		r.Post("/", matHandle.PostMaterial)
		r.Get("/{id}", matHandle.GetMaterialByID)
		r.Put("/{id}", matHandle.UpdateMaterial)
		r.Delete("/{id}", matHandle.DeleteMaterial)
	})
}
