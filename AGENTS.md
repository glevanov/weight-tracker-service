# Agent Guidelines

## Build/Lint/Test Commands

### TypeScript/Node.js (root directory)

```bash
# Install dependencies
pnpm install

# Run all lint checks
pnpm run lint

# Run specific linters
pnpm run lint:prettier    # Check code formatting
pnpm run lint:eslint      # Run ESLint on src/
pnpm run lint:tsc         # TypeScript type checking (no emit)

# Fix lint issues
pnpm exec prettier --write .      # Fix all formatting
pnpm exec eslint src --ext .ts --fix  # Fix ESLint issues

# Run tests
pnpm test                 # Run all Vitest tests
pnpm vitest run <pattern> # Run specific test(s) by pattern
pnpm vitest run src/validation/validation.test.ts  # Run single test file

# Start the application
pnpm start               # Compile TypeScript and run node dist/index.js
```

### Go (go/ directory)

```bash
cd go/

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

# Cleanup
make clean              # Remove built binary

# Docker
make docker-build       # Build Docker test image
```

## Code Style Guidelines

### TypeScript

**Formatting:**

- 2-space indentation
- LF line endings
- Use Prettier for all formatting
- Always include trailing semicolons
- Use single quotes for strings

**Imports:**

- Use ES modules (`"type": "module"` in package.json)
- Import type-only dependencies with `import type { ... }`
- Use `.js` extension for all local imports (NodeNext module resolution)
- Group imports: 1) Node built-ins 2) third-party 3) local modules
- Example:

```typescript
import { createServer, type IncomingMessage } from "node:http";
import { Router } from "node-router";
import { config } from "./config.js";
```

**Types:**

- Strict TypeScript enabled (strict: true)
- Target ESNext, use ESNext lib
- Prefer `type` aliases for object shapes
- Result type pattern: `Result<Data> = SuccessResult<Data> | ErrorResult`
- Explicit return types on exported functions

**Naming:**

- PascalCase: Types, interfaces, classes, enums
- camelCase: Variables, functions, methods, properties
- SCREAMING_SNAKE_CASE: Constants
- Custom errors: `{Name}Error` extending Error

**Error Handling:**

- Use custom Error classes for domain errors
- Return discriminated union types for results: `{ isSuccess: boolean; data?: T; error?: string }`
- Handle async errors with try/catch
- Log errors appropriately

### Go

**Formatting:**

- Standard `gofmt` formatting
- Use `goimports` for import management
- Tab indentation (handled by gofmt)

**Structure:**

- Main entry: `cmd/server/main.go`
- Internal packages: `internal/` (not importable by external packages)
- Handlers in `internal/handlers/`
- Config in `internal/config/`
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

### Git Hooks

Pre-commit runs automatically:

```bash
pnpm exec lint-staged && pnpm run lint:tsc
```

This formats staged files and runs type checking.

### Project Structure

```
/                     # TypeScript/Node.js service
├── src/
│   ├── handlers/     # HTTP route handlers
│   ├── auth/         # Authentication middleware
│   ├── validation/   # Input validation
│   ├── i18n/         # Internationalization
│   │   └── locales/  # Translation files
│   └── *.ts          # Core modules
├── go/               # Go service (separate)
│   ├── cmd/server/   # Entry point
│   ├── internal/
│   │   ├── config/   # Configuration
│   │   ├── handlers/ # HTTP route handlers
│   │   ├── i18n/     # Internationalization
│   │   │   └── locales/  # Translation files (en, ru, sv)
│   │   └── validation/ # Input validation
│   └── tests/        # Integration tests
└── package.json      # Node dependencies
```

### Key Conventions

1. Both services use the same port (3000) but are separate implementations
2. TypeScript service is the primary implementation
3. Go service is being built as a rewrite/alternative
4. Both use JSON response format with `{ isSuccess, data }` or `{ isSuccess, error }`
5. CORS configured to allow requests from `FRONTEND_URL`
6. Environment variables: `PORT`, `FRONTEND_URL`

### Dependencies

**TypeScript:**

- `node-router` - Custom router (GitHub dependency)
- `mongodb` - Database client
- `jsonwebtoken` - JWT authentication

**Go:**

- `github.com/go-chi/chi/v5` - HTTP router
- `github.com/go-chi/cors` - CORS middleware
- `github.com/testcontainers/testcontainers-go` - Integration testing
- `github.com/stretchr/testify` - Test assertions
