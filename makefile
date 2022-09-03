.PHONY: sqlc test server coverage up transfer

# SERVER
server:
	@docker compose up -d
	@make up
	@go run cmd/app/*.go
	@docker compose stop

create-account:
	@curl -v -X POST -H 'Accept: application/json' -H 'Content-Type: application/json' --data '{"owner":"Tobby","balance":777,"currency":"USD"}' http://localhost:8080/accounts

update-account:
	@curl -v -X PUT -H 'Accept: application/json' -H 'Content-Type: application/json' --data '{"id":1,"owner":"Tobby","balance":88888}' http://localhost:8080/accounts/1

transfer:
	@curl -v  POST 'localhost:8080/transfers' -H 'Accept: application/json' -H'Content-Type: application/json' --data '{"from_account_id": 1,"to_account_id": 2,"amount": 10,"currency": "CAD"}'

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
	@goose -dir ./internal/repository/migrations postgres "postgresql://user:password@localhost:5432/db?sslmode=disable" up

down:
	@goose -dir ./internal/repository/migrations  postgres "postgresql://user:password@localhost:5432/db?sslmode=disable" down

gooseinit:
	goose -dir ./repository/migrations create init sql

# MOCK
mock:
	@mockgen -source internal/repository/sqlc/store.go -destination tests/internal/repository/sqlc/mock/store.go -package repoMock -aux_files repository=internal/repository/sqlc/querier.go