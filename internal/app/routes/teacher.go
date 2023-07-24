package routes

import (
	"github.com/azka1415/crud-bootcamp-evermos/internal/app/handlers"
	"github.com/go-chi/chi"
)

func SetTeacherRoutes(r chi.Router) {
	teaHandle := handlers.NewTeacherHandler()
	r.Route("/teachers", func(r chi.Router) {
		r.Get("/", teaHandle.GetTeacher)
		r.Post("/", teaHandle.PostTeacher)
		r.Get("/{id}", teaHandle.GetTeacherByID)
		r.Put("/{id}", teaHandle.UpdateTeacher)
		r.Delete("/{id}", teaHandle.DeleteTeacher)
	})
}
