package routes

import (
	"database/sql"

	"github.com/azka1415/crud-bootcamp-evermos/internal/app/handlers"
	"github.com/azka1415/crud-bootcamp-evermos/internal/app/models"
	"github.com/azka1415/crud-bootcamp-evermos/internal/app/repository"
	"github.com/go-chi/chi"
)

func SetMaterialRoutes(r chi.Router, db *sql.DB) {
	matRepo := repository.NewMaterialRepository(db)
	matService := models.NewMaterialService(matRepo)
	matHandle := handlers.NewMaterialHandler(matService)
	r.Route("/materials", func(r chi.Router) {
		r.Get("/", matHandle.GetMaterial)
		r.Post("/", matHandle.PostMaterial)
		r.Get("/{id}", matHandle.GetMaterialByID)
		r.Put("/{id}", matHandle.UpdateMaterial)
		r.Delete("/{id}", matHandle.DeleteMaterial)
	})
}
