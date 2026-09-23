package main

import (
	"log"
	"net/http"

	"github.com/tobechi-dev/git-viz/internal/app"
)

func main() {
	a, err := app.New()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Server running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", a.Handler()); err != nil {
		log.Fatal(err)
	}
}
