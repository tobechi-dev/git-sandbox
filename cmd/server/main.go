package main

import (
	"log"
	"net/http"

	inertia "github.com/romsar/gonertia/v3"
)

func main() {
	// Load the Inertia root template
	i, err := inertia.NewFromFile("resources/views/root.html")
	if err != nil {
		log.Fatal(err)
	}

	// Wrap with Vite integration
	vi, err := inertia.NewVite(i,
		inertia.WithBuildManifest("web/dist/.vite/manifest.json"),
		inertia.WithEntryPoints("src/app.jsx"),
	)
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()

	// Serve Vite build assets at /build/*
	// Strips /build/ prefix and serves files from web/dist/
	mux.Handle("/build/",
		http.StripPrefix("/build/", http.FileServer(http.Dir("web/dist"))),
	)

	// Home page
	mux.Handle("/", vi.Middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := vi.Render(w, r, "Home", inertia.Props{
			"message": "Git Sandbox is alive 🚀",
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})))

	log.Println("Server running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
