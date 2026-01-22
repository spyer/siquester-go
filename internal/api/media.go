package api

import (
	"io"
	"net/http"
	"net/url"

	"github.com/VladimirKhil/SI/siquester-go/internal/siq"
	"github.com/go-chi/chi/v5"
)

func listMedia(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	mediaType := chi.URLParam(r, "type")
	doc, ok := store.Get(id)
	if !ok {
		respondError(w, http.StatusNotFound, "Package not found")
		return
	}
	files := doc.ListMedia(mediaType)
	if files == nil {
		files = []string{}
	}
	respondJSON(w, http.StatusOK, map[string]interface{}{
		"type":  mediaType,
		"files": files,
	})
}

func getMedia(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	mediaType := chi.URLParam(r, "type")
	name, err := url.PathUnescape(chi.URLParam(r, "name"))
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid file name")
		return
	}
	doc, ok := store.Get(id)
	if !ok {
		respondError(w, http.StatusNotFound, "Package not found")
		return
	}
	data, ok := doc.GetMedia(mediaType, name)
	if !ok {
		respondError(w, http.StatusNotFound, "Media not found")
		return
	}
	// Set content type based on media type
	contentType := getContentType(mediaType, name)
	w.Header().Set("Content-Type", contentType)
	w.Header().Set("Content-Disposition", "inline; filename=\""+name+"\"")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func uploadMedia(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	mediaType := chi.URLParam(r, "type")
	doc, ok := store.Get(id)
	if !ok {
		respondError(w, http.StatusNotFound, "Package not found")
		return
	}
	if err := r.ParseMultipartForm(100 << 20); err != nil { // 100MB max
		respondError(w, http.StatusBadRequest, "Failed to parse form")
		return
	}
	file, header, err := r.FormFile("file")
	if err != nil {
		respondError(w, http.StatusBadRequest, "No file provided")
		return
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		respondError(w, http.StatusBadRequest, "Failed to read file")
		return
	}
	name := header.Filename
	if err := doc.SetMedia(mediaType, name, data); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, map[string]string{
		"type": mediaType,
		"name": name,
	})
}

func deleteMedia(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	mediaType := chi.URLParam(r, "type")
	name, err := url.PathUnescape(chi.URLParam(r, "name"))
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid file name")
		return
	}
	doc, ok := store.Get(id)
	if !ok {
		respondError(w, http.StatusNotFound, "Package not found")
		return
	}
	if !doc.DeleteMedia(mediaType, name) {
		respondError(w, http.StatusNotFound, "Media not found")
		return
	}
	respondJSON(w, http.StatusOK, map[string]string{"status": "deleted"})
}

func getContentType(mediaType, filename string) string {
	switch mediaType {
	case siq.ContentTypeImage, siq.CollectionImages:
		// Try to determine from extension
		if len(filename) > 4 {
			switch filename[len(filename)-4:] {
			case ".png":
				return "image/png"
			case ".gif":
				return "image/gif"
			case ".svg":
				return "image/svg+xml"
			case "webp":
				return "image/webp"
			}
			if len(filename) > 5 && filename[len(filename)-5:] == ".jpeg" {
				return "image/jpeg"
			}
		}
		return "image/jpeg"
	case siq.ContentTypeAudio, siq.CollectionAudio:
		if len(filename) > 4 {
			switch filename[len(filename)-4:] {
			case ".mp3":
				return "audio/mpeg"
			case ".wav":
				return "audio/wav"
			case ".ogg":
				return "audio/ogg"
			}
		}
		return "audio/mpeg"
	case siq.ContentTypeVideo, siq.CollectionVideo:
		if len(filename) > 4 {
			switch filename[len(filename)-4:] {
			case ".mp4":
				return "video/mp4"
			case ".avi":
				return "video/x-msvideo"
			case "webm":
				return "video/webm"
			}
		}
		return "video/mp4"
	case siq.ContentTypeHTML, siq.CollectionHTML:
		return "text/html"
	default:
		return "application/octet-stream"
	}
}
