DB_PORT=5450


up:
	@docker compose up db -d
	@make up
	@go run cmd/app/*.go
	
signup:
	@curl -v -X POST -H 'Accept: application/json' -H 'Content-Type: application/json' --data '{"owner":"Tobby","username":"Tobb262","password":"123","balance":777,"currency":"USD"}' http://localhost:8080/auth/signup

signin:
	@curl -v -c cookie.txt -X POST -H 'Accept: application/json' -H 'Content-Type: application/json' --data '{"username":"Tobb262","password":"123"}' http://localhost:8080/auth/signin

update-account:
	@curl -v -b ./cookie.txt -X PUT -H 'Accept: application/json' -H 'Content-Type: application/json' --data '{"id":7,"owner":"Johny","balance":132}' http://localhost:8080/accounts/7

transfer:
	@curl -v -b  ./cookie.txt  POST 'localhost:8080/accounts/transfers' -H 'Accept: application/json' -H'Content-Type: application/json' --data '{"from_account_id": 2,"to_account_id": 1,"amount": 10,"currency": "USD"}'

get-account:
	@curl -v -b ./cookie.txt -X GET 'localhost:8080/accounts/1' 
	
delete-account:
	@curl -v -b ./cookie.txt -v -X DELETE 'localhost:8080/accounts/1' 

get-accounts:
	@curl -v -b ./cookie.txt -v -X GET 'localhost:8080/accounts?page_id=1&page_size=5' 

# TESTS
test:
	@docker compose up db -d
	@make up
	@go test -v -cover  -coverpkg=./... ./...
	@docker compose stop

testCI:
	@go test -v -cover  -coverpkg=./... ./...
	
coverage:
	@docker compose up db -d
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
	@goose -dir ./internal/repository/migrations postgres "postgresql://user:password@localhost:${DB_PORT}/db?sslmode=disable" up

down:
	@goose -dir ./internal/repository/migrations  postgres "postgresql://user:password@localhost:${DB_PORT}/db?sslmode=disable" down

gooseinit:
	goose -dir ./repository/migrations create init sql

# MOCK
mock:
	@mockgen -source internal/repository/sqlc/store.go -destination tests/internal/repository/sqlc/mock/store.go -package repoMock -aux_files repository=internal/repository/sqlc/querier.go


	# MIGRATIONS
upDoc:
	@goose -dir ./internal/repository/migrations postgres "postgresql://user:password@0.0.0.0:${DB_PORT}/db?sslmode=disable" up

downDoc:
	@goose -dir ./internal/repository/migrations  postgres "postgresql://user:password@0.0.0.0:${DB_PORT}/db?sslmode=disable" down
