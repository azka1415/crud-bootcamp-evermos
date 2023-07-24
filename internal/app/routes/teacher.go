package routes

import (
	"database/sql"

	"github.com/azka1415/crud-bootcamp-evermos/internal/app/handlers"
	"github.com/azka1415/crud-bootcamp-evermos/internal/app/models"
	"github.com/azka1415/crud-bootcamp-evermos/internal/app/repository"
	"github.com/go-chi/chi"
)

func SetTeacherRoutes(r chi.Router, db *sql.DB) {
	teaRepo := repository.NewTeacherRepository()
	teaService := models.NewTeacherService(teaRepo)
	teaHandle := handlers.NewTeacherHandler(teaService)
	r.Route("/teachers", func(r chi.Router) {
		r.Get("/", teaHandle.GetTeacher)
		r.Post("/", teaHandle.PostTeacher)
		r.Get("/{id}", teaHandle.GetTeacherByID)
		r.Put("/{id}", teaHandle.UpdateTeacher)
		r.Delete("/{id}", teaHandle.DeleteTeacher)
	})
}
