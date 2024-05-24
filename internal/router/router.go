package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/yunuscvlk/shellix-task/internal/db"
	"github.com/yunuscvlk/shellix-task/internal/handlers/sentences"
)

type Router struct {
	Mux *chi.Mux
	LastError error
}

func Initialize(db *db.DB) *Router {
	instance := new(Router)

	instance.Mux = chi.NewRouter()

	instance.Mux.Route("/sentences", func(r chi.Router) {
		// Create
		r.Post("/", func(w http.ResponseWriter, r *http.Request) {
			sentences.CreateHandler(w, r, db)
		})

		// Read
		r.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
			sentences.ReadHandler(w, r, db)
		})

		// Read All
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			sentences.ReadAllHandler(w, r, db)
		})

		// Update
		r.Patch("/{id}", func(w http.ResponseWriter, r *http.Request) {
			sentences.UpdateHandler(w, r, db)
		})
		
		// Delete
		r.Delete("/{id}", func(w http.ResponseWriter, r *http.Request) {
			sentences.DeleteHandler(w, r, db)
		})
	})

	return instance
}