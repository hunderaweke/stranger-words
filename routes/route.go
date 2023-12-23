package routes

import (
	"stranger-words/controllers"

	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(r *chi.Mux) {
	r.Get("/api/words/", controllers.GetWords)
	r.Get("/api/words/{id}", controllers.GetWord)
}
