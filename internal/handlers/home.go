package handlers

import (
	"net/http"

	inertia "github.com/romsar/gonertia/v3"
)

// Home handles the root page.
type Home struct {
	inertia *inertia.ViteInstance
}

// NewHome builds a Home handler with its dependencies.
func NewHome(i *inertia.ViteInstance) *Home {
	return &Home{inertia: i}
}

// ServeHTTP renders the Home page.
func (h *Home) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := h.inertia.Render(w, r, "Home", inertia.Props{
		"message": "Git Sandbox is alive 🚀",
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
