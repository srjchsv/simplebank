.PHONY: sqlc test server coverage up

# SERVER
server:
	@go run cmd/app/*.go

create-account:
	@curl -v -X POST -H 'Accect: application/json' -H 'Content-Type: application/json' --data '{"owner":"Tobby","currency":"USD"}' http://localhost:8080/accounts

get-account:
	@curl -v -X GET 'localhost:8080/accounts/1' 
	
delete-account:
	@curl -v -X DELETE 'localhost:8080/accounts/1' 

get-accounts:
	@curl -v -X GET 'localhost:8080/accounts?page_id=1&page_size=5' 

# TESTS
test:
	@docker compose up -d
	@make up
	@go test -v -cover  -coverpkg=./... ./...
	@docker compose stop

coverage:
	@docker compose up -d
	@make up
	@go test -coverprofile=coverage.out -coverpkg=./... ./...
	@go tool cover -html=coverage.out
	@rm coverage.out
	@docker compose stop

# SQLC
sqlc:
	@sqlc generate

# MIGRATIONS
up:
	@goose -dir ./internal/repository/migrations postgres "user=user password=password dbname=db sslmode=disable" up

down:
	@goose -dir ./internal/repository/migrations  postgres "user=user password=password dbname=db sslmode=disable" down

gooseinit:
	goose -dir ./repository/migrations create init sql

# MOCK
mock:
	@mockgen -source internal/repository/sqlc/store.go -destination tests/internal/repository/sqlc/mock/store.go -package repoMock -aux_files repository=internal/repository/sqlc/querier.go