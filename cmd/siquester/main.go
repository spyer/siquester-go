package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/spyer/siquester-go/internal/api"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	port := flag.Int("port", 8080, "HTTP server port")
	webDir := flag.String("web", "", "Path to web directory (defaults to ./web)")
	flag.Parse()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Compress(5))

	// API routes
	api.RegisterRoutes(r)

	// Serve static files
	webPath := *webDir
	if webPath == "" {
		// Try to find web directory relative to executable
		exe, err := os.Executable()
		if err == nil {
			webPath = filepath.Join(filepath.Dir(exe), "web")
		}
		if _, err := os.Stat(webPath); os.IsNotExist(err) {
			// Try current directory
			webPath = "web"
		}
	}
	fmt.Printf("Serving web files from: %s\n", webPath)
	fileServer := http.FileServer(http.Dir(webPath))
	r.Handle("/*", fileServer)

	addr := fmt.Sprintf(":%d", *port)
	fmt.Printf("SIQuester server starting on http://localhost%s\n", addr)

	if err := http.ListenAndServe(addr, r); err != nil {
		fmt.Fprintf(os.Stderr, "Server error: %v\n", err)
		os.Exit(1)
	}
}
