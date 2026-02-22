# Agent Guidelines

## Build/Lint/Test Commands
```bash
# Build
make build              # Build the binary (./server)
make run                # Run the server directly

# Testing
make test              # Run integration tests with testcontainers
make test-unit         # Run unit tests (internal packages)
make test-verbose      # Run tests with verbose output and no cache
TESTCONTAINERS_RYUK_DISABLED=true go test -v ./tests/...  # Direct test command

# Dependencies
make tidy               # Run go mod tidy
make download           # Download dependencies

# Formatting
make fmt                # Run go fmt on all packages

# Cleanup
make clean              # Remove built binary

# Docker
make docker-build       # Build Docker test image
```

## Code Style Guidelines
**Formatting:**

- Standard `gofmt` formatting
- Use `goimports` for import management
- Tab indentation (handled by gofmt)

**Structure:**

- Main entry: `cmd/server/main.go`
- Internal packages: `internal/` (not importable by external packages)
- Handlers in `internal/handlers/`
- Config in `internal/config/`
- Database in `internal/database/`
- Tests in `tests/` package

**Naming:**

- PascalCase: Exported identifiers (types, functions, constants)
- camelCase: Unexported identifiers
- Acronyms: All caps (URL, HTTP, ID)
- Test files: `*_test.go`

**Error Handling:**

- Return errors as last return value
- Use `github.com/stretchr/testify/assert` and `require` for tests
- Always check errors with `if err != nil`

**Testing:**

- Integration tests use testcontainers
- Test helper functions return `(value, cleanup)` pattern
- Use `t.Helper()` in test helpers
- Set `TESTCONTAINERS_RYUK_DISABLED=true` for test runs

### Project Structure

```
├── cmd/server/   # Entry point
├── internal/
│   ├── config/   # Configuration
│   ├── database/ # Database connection
│   ├── handlers/ # HTTP route handlers
│   ├── i18n/     # Internationalization
│   │   └── locales/  # Locale Go files (en.go, ru.go, sv.go, locale.go)
│   └── validation/ # Input validation
└── tests/        # Integration tests
```

### Key Conventions

1. Both services use the same port (3000) but are separate implementations
2. TypeScript service is the primary implementation
3. Go service is being built as a rewrite/alternative
4. Both use JSON response format with `{ isSuccess, data }` or `{ isSuccess, error }`
5. CORS configured to allow requests from `FRONTEND_URL`
6. Environment variables: `PORT`, `FRONTEND_URL`

### Dependencies

- `github.com/go-chi/chi/v5` - HTTP router
- `github.com/go-chi/cors` - CORS middleware
- `github.com/testcontainers/testcontainers-go` - Integration testing
- `github.com/stretchr/testify` - Test assertions
