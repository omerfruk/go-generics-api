# Go Generics REST API

A compact REST API that demonstrates how Go generics can remove repetitive CRUD
code while keeping resource-specific behavior explicit. The project uses Fiber
for HTTP routing, GORM for persistence, and PostgreSQL for storage.

## What it demonstrates

- Generic CRUD services and HTTP handlers
- Separate model, service, handler, and routing layers
- Resource-specific queries alongside reusable generic operations
- Automatic database migrations for users and books
- Environment-based local configuration

## Stack

- Go 1.18+
- Fiber
- GORM
- PostgreSQL

## Run locally

1. Copy the example environment file:

   ```bash
   cp .env.example .env
   ```

2. Start PostgreSQL:

   ```bash
   docker compose up -d
   ```

3. Download dependencies and start the API:

   ```bash
   go mod download
   go run .
   ```

The API listens on `http://localhost:4747` by default. PostgreSQL data is stored
in a named Docker volume so it survives container restarts.

## Configuration

| Variable | Default | Purpose |
|---|---|---|
| `APP_PORT` | `4747` | HTTP server port |
| `DB_HOST` | `localhost` | PostgreSQL host |
| `DB_PORT` | `5432` | PostgreSQL port |
| `DB_USER` | `go_generics` | Database user |
| `DB_PASSWORD` | `go_generics` | Local development password |
| `DB_NAME` | `go_generics` | Database name |
| `DB_SSLMODE` | `disable` | PostgreSQL SSL mode |

Use different credentials outside local development and never commit a populated
`.env` file.

## API routes

| Method | Route | Description |
|---|---|---|
| `GET` | `/users` | List users |
| `GET` | `/users/:id` | Get a user by ID |
| `POST` | `/users` | Create a user |
| `PUT` | `/users/:id` | Update a user |
| `DELETE` | `/users/:id` | Delete a user |
| `GET` | `/users/email/:email` | Find a user by email |
| `GET` | `/books` | List books |
| `GET` | `/books/:id` | Get a book by ID |
| `POST` | `/books` | Create a book |
| `PUT` | `/books/:id` | Update a book |
| `DELETE` | `/books/:id` | Delete a book |
| `GET` | `/books/author/:author` | Find books by author |

Example request:

```bash
curl -X POST http://localhost:4747/books \
  -H 'Content-Type: application/json' \
  -d '{"name":"The Little Go Book","author":"Karl Seguin","description":"A concise introduction to Go"}'
```

## Project structure

```text
.
├── database/  # PostgreSQL connection and migrations
├── handlers/  # Generic and resource-specific HTTP handlers
├── model/     # GORM models
├── router/    # Route registration
└── services/  # Generic CRUD and resource-specific queries
```

## Verification

```bash
go test ./...
go build ./...
```

## Scope

This is an educational architecture example, not a production-ready identity
service. Authentication, password hashing, request validation, authorization,
and production observability are intentionally outside its current scope.

For additional context, see the accompanying
[article on Go generics](https://medium.com/@omer.fruk/go-generic-ile-api-5a3594aa2763).
