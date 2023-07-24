package internal

import (
	"net/http"

	"github.com/azka1415/crud-bootcamp-evermos/db"
	"github.com/azka1415/crud-bootcamp-evermos/internal/app/middleware"
	"github.com/azka1415/crud-bootcamp-evermos/internal/app/routes"
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
	r.Use(middleware.CorsMiddleware)
	r.Use(middleware.JsonMiddleware)
	routes.SetMaterialRoutes(r)
	apiLogger.Info("Connected to database")
	defer db.Close()
	apiLogger.Info("Server starting on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
