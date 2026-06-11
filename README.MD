# Simple Bank

Simple Bank is a Go backend service for managing accounts and transferring money between them. It uses PostgreSQL for persistence, `sqlc` for type-safe query generation, and Gin for the HTTP API.

## Tech Stack

- Go 1.26.3
- Gin HTTP framework
- PostgreSQL 18 via Docker
- `pgx/v5` PostgreSQL driver and pool
- `sqlc` for generated database access code
- `golang-migrate` for schema migrations
- `viper` for environment-based configuration

## Project Structure

```text
.
├── api/             # HTTP handlers, request validation, server setup
├── db/
│   ├── migration/   # Database schema migrations
│   ├── query/       # SQL queries used by sqlc
│   ├── sqlc/        # Generated database code and store implementation
│   └── mock/        # Generated mocks for tests
├── util/            # Config loading, currency helpers, test utilities
├── app.env          # Local environment variables
├── main.go          # Application entry point
├── Makefile         # Common development commands
└── sqlc.yaml        # sqlc configuration
```

## Prerequisites

Install these tools before running the project locally:

- Go
- Docker
- `sqlc`
- `golang-migrate`
- `mockgen` if you want to regenerate mocks

## Configuration

The app reads configuration from `app.env`.

```env
DB_DRIVER=postgres
DB_SOURCE=postgresql://root:secret@localhost:5433/simple_bank?sslmode=disable
SERVER_ADDRESS=0.0.0.0:8080
```

## Setup

Start PostgreSQL in Docker:

```sh
make postgres
```

Create the development database:

```sh
make createdb
```

Run database migrations:

```sh
make migrateup
```

Generate Go code from SQL queries:

```sh
make sqlc
```

## Run The Server

Start the API server:

```sh
make server
```

The default local server address is:

```text
http://localhost:8080
```

## API Endpoints

The server currently exposes these routes:

- `POST /accounts` creates an account
- `GET /accounts/:id` fetches an account by ID
- `GET /accounts?page_id=1&page_size=5` lists accounts with pagination
- `POST /transfers` creates a transfer between two accounts

Request validation includes:

- positive account IDs
- positive transfer amounts
- supported currencies only: `USD`, `EUR`, `CAD`

Example create account request:

```json
{
  "owner": "alice",
  "currency": "USD"
}
```

Example transfer request:

```json
{
  "from_account_id": 1,
  "to_account_id": 2,
  "amount": 100,
  "currency": "USD"
}
```

## Useful Commands

Run tests:

```sh
make test
```

Generate mocks:

```sh
make mock
```

Rollback migrations:

```sh
make migratedown
```

Drop the database:

```sh
make dropdb
```

## Database

The schema includes:

- `accounts` for owners, balances, and currencies
- `entries` for account balance changes
- `transfers` for money movement between accounts

The default local database connection used by the Makefile is:

```text
postgresql://root:secret@localhost:5433/simple_bank?sslmode=disable
```
