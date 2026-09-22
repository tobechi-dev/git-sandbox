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

	mux := http.NewServeMux()

	// Home page — renders the "Home" React component
	mux.Handle("/", i.Middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := i.Render(w, r, "Home", inertia.Props{
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
