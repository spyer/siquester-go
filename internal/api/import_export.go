package api

import (
	"encoding/xml"
	"io"
	"net/http"

	"github.com/VladimirKhil/SI/siquester-go/internal/siq"
	"github.com/go-chi/chi/v5"
	"gopkg.in/yaml.v3"
)

func importXML(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		respondError(w, http.StatusBadRequest, "Failed to read request body")
		return
	}
	var pkg siq.Package
	if err := xml.Unmarshal(data, &pkg); err != nil {
		respondError(w, http.StatusBadRequest, "Failed to parse XML: "+err.Error())
		return
	}
	doc := &siq.Document{
		Package: &pkg,
		Images:  make(map[string][]byte),
		Audio:   make(map[string][]byte),
		Video:   make(map[string][]byte),
		HTML:    make(map[string][]byte),
	}
	id := store.Add(doc)
	respondJSON(w, http.StatusCreated, PackageResponse{
		ID:      id,
		Package: doc.Package,
	})
}

func importYAML(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		respondError(w, http.StatusBadRequest, "Failed to read request body")
		return
	}
	var pkg siq.Package
	if err := yaml.Unmarshal(data, &pkg); err != nil {
		respondError(w, http.StatusBadRequest, "Failed to parse YAML: "+err.Error())
		return
	}
	// Set defaults
	if pkg.Version == 0 {
		pkg.Version = siq.PackageVersion
	}
	doc := &siq.Document{
		Package: &pkg,
		Images:  make(map[string][]byte),
		Audio:   make(map[string][]byte),
		Video:   make(map[string][]byte),
		HTML:    make(map[string][]byte),
	}
	id := store.Add(doc)
	respondJSON(w, http.StatusCreated, PackageResponse{
		ID:      id,
		Package: doc.Package,
	})
}

func exportXML(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	doc, ok := store.Get(id)
	if !ok {
		respondError(w, http.StatusNotFound, "Package not found")
		return
	}
	w.Header().Set("Content-Type", "application/xml")
	w.Header().Set("Content-Disposition", "attachment; filename=\""+doc.Package.Name+".xml\"")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(xml.Header))
	encoder := xml.NewEncoder(w)
	encoder.Indent("", "  ")
	encoder.Encode(doc.Package)
}

func exportYAML(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	doc, ok := store.Get(id)
	if !ok {
		respondError(w, http.StatusNotFound, "Package not found")
		return
	}
	w.Header().Set("Content-Type", "application/x-yaml")
	w.Header().Set("Content-Disposition", "attachment; filename=\""+doc.Package.Name+".yaml\"")
	w.WriteHeader(http.StatusOK)
	encoder := yaml.NewEncoder(w)
	encoder.SetIndent(2)
	encoder.Encode(doc.Package)
}

func exportSIQ(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	doc, ok := store.Get(id)
	if !ok {
		respondError(w, http.StatusNotFound, "Package not found")
		return
	}
	data, err := doc.ToBytes()
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", "attachment; filename=\""+doc.Package.Name+".siq\"")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
