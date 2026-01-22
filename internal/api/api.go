package api

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// RegisterRoutes registers all API routes.
func RegisterRoutes(r chi.Router) {
	r.Route("/api", func(r chi.Router) {
		// Package operations
		r.Post("/packages", createPackage)
		r.Get("/packages", listPackages)
		r.Get("/packages/{id}", getPackage)
		r.Put("/packages/{id}", updatePackage)
		r.Delete("/packages/{id}", deletePackage)
		r.Post("/packages/{id}/save", savePackage)
		r.Post("/packages/open", openPackage)
		// Round operations
		r.Post("/packages/{id}/rounds", addRound)
		r.Put("/packages/{id}/rounds/{roundIndex}", updateRound)
		r.Delete("/packages/{id}/rounds/{roundIndex}", deleteRound)
		r.Post("/packages/{id}/rounds/reorder", reorderRounds)
		// Theme operations
		r.Post("/packages/{id}/rounds/{roundIndex}/themes", addTheme)
		r.Put("/packages/{id}/rounds/{roundIndex}/themes/{themeIndex}", updateTheme)
		r.Delete("/packages/{id}/rounds/{roundIndex}/themes/{themeIndex}", deleteTheme)
		r.Post("/packages/{id}/themes/move", moveTheme)
		// Question operations
		r.Post("/packages/{id}/rounds/{roundIndex}/themes/{themeIndex}/questions", addQuestion)
		r.Put("/packages/{id}/rounds/{roundIndex}/themes/{themeIndex}/questions/{questionIndex}", updateQuestion)
		r.Delete("/packages/{id}/rounds/{roundIndex}/themes/{themeIndex}/questions/{questionIndex}", deleteQuestion)
		// Media operations
		r.Get("/packages/{id}/media/{type}", listMedia)
		r.Get("/packages/{id}/media/{type}/{name}", getMedia)
		r.Post("/packages/{id}/media/{type}", uploadMedia)
		r.Delete("/packages/{id}/media/{type}/{name}", deleteMedia)
		// Import/Export
		r.Post("/import/xml", importXML)
		r.Post("/import/yaml", importYAML)
		r.Get("/export/{id}/xml", exportXML)
		r.Get("/export/{id}/yaml", exportYAML)
		r.Get("/export/{id}/siq", exportSIQ)
	})
}

// JSON response helpers
func respondJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

func respondError(w http.ResponseWriter, status int, message string) {
	respondJSON(w, status, map[string]string{"error": message})
}
