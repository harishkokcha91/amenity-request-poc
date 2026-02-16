# Amenity Request POC

Small backend POC implementing an Amenity Request feature using Go, Gin and PostgreSQL.

Contents
- API: controllers, services, repositories for amenity requests
- DB: Postgres migration in `migrations/001_create_amenity_requests.sql`
- Dev helpers: `docker-compose.yml`, `Makefile`, and `all_api.rest` (REST Client)

Prerequisitesm
- Go 1.20+
- Docker and Docker Compose (or Docker Desktop)
- (Optional) VS Code + REST Client extension to open `all_api.rest`

Quick start (recommended)

1. Create `.env` from the example (edit if you need custom credentials):

```bash
cp .env.example .env
```

2. Start Postgres with Docker Compose (this will run SQL files in `migrations/` on first startup):

```bash
make up
```

3. Start the app (waits for Postgres, applies migration, runs server):

```bash
make start
```

The server listens on :8080 by default. You can override `PORT` and `DATABASE_URL` via environment variables.

Useful Makefile targets
- `make up` — starts Postgres (creates `.env` from `.env.example` if missing)
- `make start` — starts Postgres, waits for readiness, applies migration, runs the Go server (foreground)
- `make run` — runs Go server using `DATABASE_URL` from `.env`
- `make migrate` — applies the SQL migration inside the Postgres container (idempotent)
- `make down` — stops containers and removes volumes
- `make test` — runs `go test ./...`

Environment
- `.env.example` contains default values:

```
POSTGRES_USER=postgres
POSTGRES_PASSWORD=password
POSTGRES_DB=amenitydb
DATABASE_URL=postgres://postgres:password@localhost:5432/amenitydb?sslmode=disable
```

API endpoints
- POST /api/v1/reservations/:id/amenity-requests — create a request
- GET /api/v1/reservations/:id/amenity-requests — list requests for a reservation
- GET /api/v1/amenity-requests/:requestId — fetch a request by id
- PATCH /api/v1/amenity-requests/:requestId/status — update status (PENDING, IN_PROGRESS, COMPLETED, CANCELLED)
- GET /api/v1/reservations — list distinct reservation ids that have amenity requests
- GET /api/v1/healthcheck — simple health endpoint

Sample curl (create):

```bash
curl -X POST http://localhost:8080/api/v1/reservations/RES123/amenity-requests \
  -H 'Content-Type: application/json' \
  -d '{"guest_id":"G1","property_id":"P1","amenity_type":"TOWELS","quantity":2}'
```

REST client
- `all_api.rest` contains ready-to-run requests (health, create, list, get-by-id, update status, auth example). Open with VS Code REST Client and click "Send Request".

Authentication
- The repo includes REST examples that send an `Authorization: Bearer <token>` header, but the server does not enforce JWT validation yet. If you want token-based auth, I can add middleware that validates JWTs (HS256 or RS256) and protects routes.

Testing
- Run unit tests with:

```bash
make test
```

Development notes and next steps
- Reservation validation: currently the service assumes a reservation exists and is checked-in; if you have a reservations table we can validate this on Create.
- completed_at: when status becomes COMPLETED we currently don't set `completed_at` — we can update the repository to set this timestamp.
- Docker: we provide Postgres via Docker Compose; if you prefer the app inside Docker too I can add a `Dockerfile` and Compose service for the Go app.

Contact / Support
If you want additional features (auth, reservation model, more tests, Dockerfile), tell me which and I'll implement them.
