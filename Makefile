postgres:
	docker run --name docker-db-1 -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -d postgres

createdb:
	docker exec -it docker-db-1 createdb --username=postgres --owner=postgres simple_bank

dropdb:
	docker exec -it docker-db-1 dropdb simple_banki

.PHONY: postgres createdb dropdb