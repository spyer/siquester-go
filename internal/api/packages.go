package api

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/VladimirKhil/SI/siquester-go/internal/siq"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

// CreatePackageRequest represents a request to create a new package.
type CreatePackageRequest struct {
	Name   string `json:"name"`
	Author string `json:"author"`
}

// PackageResponse represents a package in API responses.
type PackageResponse struct {
	ID      string       `json:"id"`
	Package *siq.Package `json:"package"`
}

func createPackage(w http.ResponseWriter, r *http.Request) {
	var req CreatePackageRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request body")
		return
	}
	if req.Name == "" {
		req.Name = "New Package"
	}
	if req.Author == "" {
		req.Author = "Unknown"
	}
	doc := siq.NewDocument(req.Name, req.Author)
	doc.Package.ID = uuid.New().String()
	doc.Package.Date = time.Now().Format("02.01.2006")
	id := store.Add(doc)
	respondJSON(w, http.StatusCreated, PackageResponse{
		ID:      id,
		Package: doc.Package,
	})
}

func listPackages(w http.ResponseWriter, r *http.Request) {
	infos := store.List()
	respondJSON(w, http.StatusOK, infos)
}

func getPackage(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	doc, ok := store.Get(id)
	if !ok {
		respondError(w, http.StatusNotFound, "Package not found")
		return
	}
	respondJSON(w, http.StatusOK, PackageResponse{
		ID:      id,
		Package: doc.Package,
	})
}

func updatePackage(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	doc, ok := store.Get(id)
	if !ok {
		respondError(w, http.StatusNotFound, "Package not found")
		return
	}
	var pkg siq.Package
	if err := json.NewDecoder(r.Body).Decode(&pkg); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request body")
		return
	}
	// Update package fields
	doc.Package.Name = pkg.Name
	doc.Package.Restriction = pkg.Restriction
	doc.Package.Date = pkg.Date
	doc.Package.Publisher = pkg.Publisher
	doc.Package.ContactURI = pkg.ContactURI
	doc.Package.Difficulty = pkg.Difficulty
	doc.Package.Logo = pkg.Logo
	doc.Package.Language = pkg.Language
	doc.Package.Tags = pkg.Tags
	doc.Package.Info = pkg.Info
	doc.Package.Rounds = pkg.Rounds
	doc.Package.Global = pkg.Global
	respondJSON(w, http.StatusOK, PackageResponse{
		ID:      id,
		Package: doc.Package,
	})
}

func deletePackage(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if !store.Delete(id) {
		respondError(w, http.StatusNotFound, "Package not found")
		return
	}
	respondJSON(w, http.StatusOK, map[string]string{"status": "deleted"})
}

func savePackage(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	doc, ok := store.Get(id)
	if !ok {
		respondError(w, http.StatusNotFound, "Package not found")
		return
	}
	var req struct {
		Path string `json:"path"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request body")
		return
	}
	path := req.Path
	if path == "" {
		path = doc.FilePath
	}
	if path == "" {
		respondError(w, http.StatusBadRequest, "No file path specified")
		return
	}
	if err := doc.SaveToFile(path); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, map[string]string{
		"status": "saved",
		"path":   path,
	})
}

func openPackage(w http.ResponseWriter, r *http.Request) {
	// Check if it's a file upload or a path request
	contentType := r.Header.Get("Content-Type")
	if contentType == "application/json" {
		// Path-based open
		var req struct {
			Path string `json:"path"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			respondError(w, http.StatusBadRequest, "Invalid request body")
			return
		}
		doc, err := siq.LoadFromFile(req.Path)
		if err != nil {
			respondError(w, http.StatusBadRequest, err.Error())
			return
		}
		id := store.Add(doc)
		respondJSON(w, http.StatusOK, PackageResponse{
			ID:      id,
			Package: doc.Package,
		})
	} else {
		// File upload
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
		doc, err := siq.LoadFromBytes(data)
		if err != nil {
			respondError(w, http.StatusBadRequest, err.Error())
			return
		}
		doc.FilePath = header.Filename
		id := store.Add(doc)
		respondJSON(w, http.StatusOK, PackageResponse{
			ID:      id,
			Package: doc.Package,
		})
	}
}

// Round operations
func addRound(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	doc, ok := store.Get(id)
	if !ok {
		respondError(w, http.StatusNotFound, "Package not found")
		return
	}
	var round siq.Round
	if err := json.NewDecoder(r.Body).Decode(&round); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request body")
		return
	}
	if round.Name == "" {
		round.Name = "New Round"
	}
	if round.Type == "" {
		round.Type = siq.RoundTypeStandard
	}
	if round.Themes == nil {
		round.Themes = make([]*siq.Theme, 0)
	}
	doc.Package.Rounds = append(doc.Package.Rounds, &round)
	respondJSON(w, http.StatusCreated, &round)
}

func updateRound(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	roundIndex, err := strconv.Atoi(chi.URLParam(r, "roundIndex"))
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid round index")
		return
	}
	doc, ok := store.Get(id)
	if !ok {
		respondError(w, http.StatusNotFound, "Package not found")
		return
	}
	if roundIndex < 0 || roundIndex >= len(doc.Package.Rounds) {
		respondError(w, http.StatusNotFound, "Round not found")
		return
	}
	var round siq.Round
	if err := json.NewDecoder(r.Body).Decode(&round); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request body")
		return
	}
	doc.Package.Rounds[roundIndex].Name = round.Name
	doc.Package.Rounds[roundIndex].Type = round.Type
	doc.Package.Rounds[roundIndex].Info = round.Info
	respondJSON(w, http.StatusOK, doc.Package.Rounds[roundIndex])
}

func deleteRound(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	roundIndex, err := strconv.Atoi(chi.URLParam(r, "roundIndex"))
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid round index")
		return
	}
	doc, ok := store.Get(id)
	if !ok {
		respondError(w, http.StatusNotFound, "Package not found")
		return
	}
	if roundIndex < 0 || roundIndex >= len(doc.Package.Rounds) {
		respondError(w, http.StatusNotFound, "Round not found")
		return
	}
	doc.Package.Rounds = append(doc.Package.Rounds[:roundIndex], doc.Package.Rounds[roundIndex+1:]...)
	respondJSON(w, http.StatusOK, map[string]string{"status": "deleted"})
}

func reorderRounds(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	doc, ok := store.Get(id)
	if !ok {
		respondError(w, http.StatusNotFound, "Package not found")
		return
	}
	var req struct {
		From int `json:"from"`
		To   int `json:"to"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request body")
		return
	}
	rounds := doc.Package.Rounds
	if req.From < 0 || req.From >= len(rounds) || req.To < 0 || req.To >= len(rounds) {
		respondError(w, http.StatusBadRequest, "Invalid round indices")
		return
	}
	// Move round from position 'from' to position 'to'
	round := rounds[req.From]
	// Remove from old position
	rounds = append(rounds[:req.From], rounds[req.From+1:]...)
	// Insert at new position
	newRounds := make([]*siq.Round, 0, len(rounds)+1)
	newRounds = append(newRounds, rounds[:req.To]...)
	newRounds = append(newRounds, round)
	newRounds = append(newRounds, rounds[req.To:]...)
	doc.Package.Rounds = newRounds
	respondJSON(w, http.StatusOK, map[string]string{"status": "reordered"})
}

func moveTheme(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	doc, ok := store.Get(id)
	if !ok {
		respondError(w, http.StatusNotFound, "Package not found")
		return
	}
	var req struct {
		FromRound int `json:"fromRound"`
		FromTheme int `json:"fromTheme"`
		ToRound   int `json:"toRound"`
		ToTheme   int `json:"toTheme"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request body")
		return
	}
	rounds := doc.Package.Rounds
	// Validate indices
	if req.FromRound < 0 || req.FromRound >= len(rounds) ||
		req.ToRound < 0 || req.ToRound >= len(rounds) {
		respondError(w, http.StatusBadRequest, "Invalid round indices")
		return
	}
	fromThemes := rounds[req.FromRound].Themes
	if req.FromTheme < 0 || req.FromTheme >= len(fromThemes) {
		respondError(w, http.StatusBadRequest, "Invalid source theme index")
		return
	}
	// Get the theme to move
	theme := fromThemes[req.FromTheme]
	// Remove from source
	rounds[req.FromRound].Themes = append(fromThemes[:req.FromTheme], fromThemes[req.FromTheme+1:]...)
	// Ensure target round has themes slice
	if rounds[req.ToRound].Themes == nil {
		rounds[req.ToRound].Themes = make([]*siq.Theme, 0)
	}
	toThemes := rounds[req.ToRound].Themes
	// Adjust target index if needed (when moving within same round to a later position)
	targetIdx := req.ToTheme
	if targetIdx > len(toThemes) {
		targetIdx = len(toThemes)
	}
	// Insert at target position
	newThemes := make([]*siq.Theme, 0, len(toThemes)+1)
	newThemes = append(newThemes, toThemes[:targetIdx]...)
	newThemes = append(newThemes, theme)
	newThemes = append(newThemes, toThemes[targetIdx:]...)
	rounds[req.ToRound].Themes = newThemes
	respondJSON(w, http.StatusOK, map[string]string{"status": "moved"})
}

// Theme operations
func addTheme(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	roundIndex, err := strconv.Atoi(chi.URLParam(r, "roundIndex"))
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid round index")
		return
	}
	doc, ok := store.Get(id)
	if !ok {
		respondError(w, http.StatusNotFound, "Package not found")
		return
	}
	if roundIndex < 0 || roundIndex >= len(doc.Package.Rounds) {
		respondError(w, http.StatusNotFound, "Round not found")
		return
	}
	var theme siq.Theme
	if err := json.NewDecoder(r.Body).Decode(&theme); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request body")
		return
	}
	if theme.Name == "" {
		theme.Name = "New Theme"
	}
	if theme.Questions == nil {
		theme.Questions = make([]*siq.Question, 0)
	}
	if doc.Package.Rounds[roundIndex].Themes == nil {
		doc.Package.Rounds[roundIndex].Themes = make([]*siq.Theme, 0)
	}
	doc.Package.Rounds[roundIndex].Themes = append(doc.Package.Rounds[roundIndex].Themes, &theme)
	respondJSON(w, http.StatusCreated, &theme)
}

func updateTheme(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	roundIndex, err := strconv.Atoi(chi.URLParam(r, "roundIndex"))
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid round index")
		return
	}
	themeIndex, err := strconv.Atoi(chi.URLParam(r, "themeIndex"))
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid theme index")
		return
	}
	doc, ok := store.Get(id)
	if !ok {
		respondError(w, http.StatusNotFound, "Package not found")
		return
	}
	if roundIndex < 0 || roundIndex >= len(doc.Package.Rounds) {
		respondError(w, http.StatusNotFound, "Round not found")
		return
	}
	round := doc.Package.Rounds[roundIndex]
	if themeIndex < 0 || themeIndex >= len(round.Themes) {
		respondError(w, http.StatusNotFound, "Theme not found")
		return
	}
	var theme siq.Theme
	if err := json.NewDecoder(r.Body).Decode(&theme); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request body")
		return
	}
	round.Themes[themeIndex].Name = theme.Name
	round.Themes[themeIndex].Info = theme.Info
	respondJSON(w, http.StatusOK, round.Themes[themeIndex])
}

func deleteTheme(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	roundIndex, err := strconv.Atoi(chi.URLParam(r, "roundIndex"))
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid round index")
		return
	}
	themeIndex, err := strconv.Atoi(chi.URLParam(r, "themeIndex"))
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid theme index")
		return
	}
	doc, ok := store.Get(id)
	if !ok {
		respondError(w, http.StatusNotFound, "Package not found")
		return
	}
	if roundIndex < 0 || roundIndex >= len(doc.Package.Rounds) {
		respondError(w, http.StatusNotFound, "Round not found")
		return
	}
	round := doc.Package.Rounds[roundIndex]
	if themeIndex < 0 || themeIndex >= len(round.Themes) {
		respondError(w, http.StatusNotFound, "Theme not found")
		return
	}
	round.Themes = append(round.Themes[:themeIndex], round.Themes[themeIndex+1:]...)
	respondJSON(w, http.StatusOK, map[string]string{"status": "deleted"})
}

// Question operations
func addQuestion(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	roundIndex, err := strconv.Atoi(chi.URLParam(r, "roundIndex"))
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid round index")
		return
	}
	themeIndex, err := strconv.Atoi(chi.URLParam(r, "themeIndex"))
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid theme index")
		return
	}
	doc, ok := store.Get(id)
	if !ok {
		respondError(w, http.StatusNotFound, "Package not found")
		return
	}
	if roundIndex < 0 || roundIndex >= len(doc.Package.Rounds) {
		respondError(w, http.StatusNotFound, "Round not found")
		return
	}
	round := doc.Package.Rounds[roundIndex]
	if themeIndex < 0 || themeIndex >= len(round.Themes) {
		respondError(w, http.StatusNotFound, "Theme not found")
		return
	}
	var question siq.Question
	if err := json.NewDecoder(r.Body).Decode(&question); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request body")
		return
	}
	if question.Price == 0 {
		// Auto-calculate price based on position
		numQuestions := 0
		if round.Themes[themeIndex].Questions != nil {
			numQuestions = len(round.Themes[themeIndex].Questions)
		}
		question.Price = (numQuestions + 1) * 100
	}
	if len(question.Right) == 0 {
		question.Right = []string{""}
	}
	if round.Themes[themeIndex].Questions == nil {
		round.Themes[themeIndex].Questions = make([]*siq.Question, 0)
	}
	round.Themes[themeIndex].Questions = append(round.Themes[themeIndex].Questions, &question)
	respondJSON(w, http.StatusCreated, &question)
}

func updateQuestion(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	roundIndex, err := strconv.Atoi(chi.URLParam(r, "roundIndex"))
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid round index")
		return
	}
	themeIndex, err := strconv.Atoi(chi.URLParam(r, "themeIndex"))
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid theme index")
		return
	}
	questionIndex, err := strconv.Atoi(chi.URLParam(r, "questionIndex"))
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid question index")
		return
	}
	doc, ok := store.Get(id)
	if !ok {
		respondError(w, http.StatusNotFound, "Package not found")
		return
	}
	if roundIndex < 0 || roundIndex >= len(doc.Package.Rounds) {
		respondError(w, http.StatusNotFound, "Round not found")
		return
	}
	round := doc.Package.Rounds[roundIndex]
	if themeIndex < 0 || themeIndex >= len(round.Themes) {
		respondError(w, http.StatusNotFound, "Theme not found")
		return
	}
	theme := round.Themes[themeIndex]
	if questionIndex < 0 || questionIndex >= len(theme.Questions) {
		respondError(w, http.StatusNotFound, "Question not found")
		return
	}
	var question siq.Question
	if err := json.NewDecoder(r.Body).Decode(&question); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request body")
		return
	}
	// Update question
	theme.Questions[questionIndex] = &question
	respondJSON(w, http.StatusOK, theme.Questions[questionIndex])
}

func deleteQuestion(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	roundIndex, err := strconv.Atoi(chi.URLParam(r, "roundIndex"))
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid round index")
		return
	}
	themeIndex, err := strconv.Atoi(chi.URLParam(r, "themeIndex"))
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid theme index")
		return
	}
	questionIndex, err := strconv.Atoi(chi.URLParam(r, "questionIndex"))
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid question index")
		return
	}
	doc, ok := store.Get(id)
	if !ok {
		respondError(w, http.StatusNotFound, "Package not found")
		return
	}
	if roundIndex < 0 || roundIndex >= len(doc.Package.Rounds) {
		respondError(w, http.StatusNotFound, "Round not found")
		return
	}
	round := doc.Package.Rounds[roundIndex]
	if themeIndex < 0 || themeIndex >= len(round.Themes) {
		respondError(w, http.StatusNotFound, "Theme not found")
		return
	}
	theme := round.Themes[themeIndex]
	if questionIndex < 0 || questionIndex >= len(theme.Questions) {
		respondError(w, http.StatusNotFound, "Question not found")
		return
	}
	theme.Questions = append(theme.Questions[:questionIndex], theme.Questions[questionIndex+1:]...)
	respondJSON(w, http.StatusOK, map[string]string{"status": "deleted"})
}
