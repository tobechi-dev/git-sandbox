# Git Sandbox

An interactive platform for learning Git commands visually. Type a Git command and see a live animated commit graph that shows what actually happens under the hood.

## Stack

- Go
- Inertia.js (`gonertia`)
- React
- Vite
- GitGraph.js / D3.js
- Tailwind CSS

## Status

Early scaffold — building the core Go + Inertia setup first.

## Project Structure

```

git-viz/
├── cmd/
│   └── server/           # Go entry point
├── internal/
│   ├── handlers/         # HTTP handlers
│   └── git/              # Git command simulation logic
├── resources/
│   └── views/
│       └── root.html     # Inertia root template (Go renders this)
├── web/                  # React + Vite frontend
├── go.mod
├── go.sum
├── LICENSE
└── README.md

```

### Folder Overview

| Folder               | Purpose                                         |
| -------------------- | ----------------------------------------------- |
| `cmd/server/`        | Go entry point — starts the HTTP server         |
| `internal/handlers/` | HTTP handlers (business logic)                  |
| `internal/git/`      | Git command parsing and simulation              |
| `resources/views/`   | HTML templates served by Go                     |
| `web/`               | Frontend source, built by Vite into `web/dist/` |

## How It Works

1. Browser hits a Go route.
2. Go renders `resources/views/root.html` (the Inertia shell).
3. The shell boots React, which takes over navigation via Inertia.
4. Subsequent page transitions are JSON responses — no full page reloads.

## Roadmap

- [x] Go server with basic routing
- [x] Inertia adapter wired up
- [ ] Root template and React frontend
- [ ] Terminal input for Git commands
- [ ] Commit graph visualization
- [ ] Command parser: `init`, `commit`, `branch`, `checkout`, `merge`
- [ ] Animated graph transitions
- [ ] Lesson / challenge mode
- [ ] Deploy to Vercel

## Getting Started

```bash
# Run the Go server
go run ./cmd/server

Server runs on http://localhost:8080

## Known Issue
Known Issue
· Termux (Android): port scanning tools (ss, lsof, fuser, netstat) don't work due to Android's /proc restrictions. If port 8080 is stuck, force-stop the Termux app or reboot.

## License

MIT
```
