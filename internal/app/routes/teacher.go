package routes

import (
	"github.com/azka1415/crud-bootcamp-evermos/internal/app/handlers"
	"github.com/go-chi/chi"
)

func SetTeacherRoutes(r chi.Router) {
	r.Route("/teachers", func(r chi.Router) {
		r.Get("/", handlers.HandleGetTeacher)

	})
}
