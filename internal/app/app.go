package app

import (
	"net/http"

	inertia "github.com/romsar/gonertia/v3"

	"github.com/tobechi-dev/git-viz/internal/handlers"
)

// App holds the wired-up application: routers, middleware, dependencies.
type App struct {
	handler http.Handler
}

// New builds the application with all its dependencies and routes.
func New() (*App, error) {
	// Infrastructure: Inertia + Vite
	i, err := inertia.NewFromFile("resources/views/root.html")
	if err != nil {
		return nil, err
	}

	vi, err := inertia.NewVite(i,
		inertia.WithBuildManifest("web/dist/.vite/manifest.json"),
		inertia.WithEntryPoints("src/app.jsx"),
		inertia.WithHotFile("web/.vite/hot"),
	)
	if err != nil {
		return nil, err
	}

	// Router
	mux := http.NewServeMux()

	// Static: Vite build output served at /build/*
	mux.Handle("/build/",
		http.StripPrefix("/build/", http.FileServer(http.Dir("web/dist"))),
	)

	// Pages
	home := handlers.NewHome(vi)
	mux.Handle("/", vi.Middleware(home))

	return &App{handler: mux}, nil
}

// Handler returns the root HTTP handler.
func (a *App) Handler() http.Handler {
	return a.handler
}
