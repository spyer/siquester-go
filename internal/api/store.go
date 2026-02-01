package api

import (
	"sync"

	"github.com/spyer/siquester-go/internal/siq"
	"github.com/google/uuid"
)

// DocumentStore manages open documents in memory.
type DocumentStore struct {
	mu        sync.RWMutex
	documents map[string]*siq.Document
}

var store = &DocumentStore{
	documents: make(map[string]*siq.Document),
}

// Add adds a document to the store and returns its ID.
func (s *DocumentStore) Add(doc *siq.Document) string {
	s.mu.Lock()
	defer s.mu.Unlock()
	id := uuid.New().String()
	s.documents[id] = doc
	return id
}

// Get retrieves a document by ID.
func (s *DocumentStore) Get(id string) (*siq.Document, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	doc, ok := s.documents[id]
	return doc, ok
}

// Delete removes a document from the store.
func (s *DocumentStore) Delete(id string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.documents[id]; ok {
		delete(s.documents, id)
		return true
	}
	return false
}

// List returns all document IDs and basic info.
func (s *DocumentStore) List() []DocumentInfo {
	s.mu.RLock()
	defer s.mu.RUnlock()
	infos := make([]DocumentInfo, 0, len(s.documents))
	for id, doc := range s.documents {
		infos = append(infos, DocumentInfo{
			ID:       id,
			Name:     doc.Package.Name,
			FilePath: doc.FilePath,
		})
	}
	return infos
}

// DocumentInfo contains basic document information.
type DocumentInfo struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	FilePath string `json:"filePath,omitempty"`
}
