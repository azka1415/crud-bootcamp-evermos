package internal

import (
	"net/http"

	"github.com/azka1415/crud-bootcamp-evermos/db"
	"github.com/azka1415/crud-bootcamp-evermos/internal/app/handlers"
	"github.com/azka1415/crud-bootcamp-evermos/internal/app/middleware"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func API() {
	apiLogger := log.WithFields(log.Fields{"Server": "Info"})
	db, err := db.GetDB()
	if err != nil {
		apiLogger.Fatal("Error opening database:", err)
	}

	r := chi.NewRouter()
	r.Use(middleware.JsonMiddleware)
	r.Get("/materials", handlers.HandleGetMaterial)
	r.Post("/materials", handlers.HandlePostMaterial)
	r.Put("/materials", handlers.HandleUpdateMaterial)
	r.Delete("/materials", handlers.HandleDeleteMaterial)
	apiLogger.Info("Connected to database")
	defer db.Close()
	apiLogger.Info("Server starting on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
