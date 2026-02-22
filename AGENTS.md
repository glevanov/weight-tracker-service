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
- Logger in `internal/logger/`
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
│   ├── auth/       # Authentication and middleware
│   ├── config/     # Configuration
│   ├── database/   # Database connection
│   ├── handlers/   # HTTP route handlers
│   ├── i18n/       # Internationalization
│   │   └── locales/    # Locale Go files (en.go, ru.go, sv.go, locale.go)
│   ├── logger/     # Structured logging (slog)
│   └── validation/ # Input validation
└── tests/          # Integration tests
```
