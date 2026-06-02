# Simple Bank

Simple Bank is a small Go project for modeling bank accounts, ledger entries, and money transfers with PostgreSQL. It uses database migrations for schema management and sqlc to generate type-safe Go code from SQL queries.

## Tech Stack

- Go 1.26.3
- PostgreSQL 18 via Docker
- pgx/v5 PostgreSQL driver
- sqlc for generated database access code
- golang-migrate for database migrations

## Project Structure

```text
.
├── db/
│   ├── migration/   # Database schema migrations
│   ├── query/       # SQL queries used by sqlc
│   └── sqlc/        # Generated Go database code
├── Makefile         # Common development commands
├── go.mod           # Go module definition
└── sqlc.yaml        # sqlc configuration
```

## Prerequisites

Install the following tools before running the project locally:

- Go
- Docker
- sqlc
- golang-migrate

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

## Useful Commands

Run tests:

```sh
go test ./...
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

The initial schema includes:

- `accounts` for account owners, balances, and currencies
- `entries` for account balance changes
- `transfers` for money movement between accounts

The default local database connection used by the Makefile is:

```text
postgresql://root:secret@localhost:5433/simple_bank?sslmode=disable
```
