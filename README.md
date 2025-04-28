# Social Backend API (Go + Chi)

This is a **production-grade backend server** built with **Go** and the **Chi router**, following clean architecture principles.

It is designed to be:
- Fast
- Scalable
- Easy to extend
- Beginner and professional friendly

---

## ✨ Features

- Go 1.23+
- HTTP server with custom configuration
- `chi` for structured and powerful routing
- Middlewares:
  - `RequestID` for request tracing
  - `RealIP` for true client IP detection
  - `Logger` for request/response logging
  - `Recoverer` for panic recovery
- Health check endpoint (`/v1/health`)
- Clean project structure
- Hot-reloading during development with `air`
- Environment-agnostic (local, staging, production ready)

---

## 📂 Project Structure

```plaintext
.
├── cmd/
│   └── api/
│       ├── api.go        # Router setup, server config
│       ├── health.go     # Health check handler
│       └── main.go       # Entry point
├── internal/
│   └── (internal packages)
├── migrate/
│   └── (database migrations )
├── scripts/
│   └── (helper scripts)
├── docs/
│   └── (API documentation)
├── go.mod
├── go.sum
└── .air.toml             # Air config for hot reload
