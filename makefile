.PHONY: sqlc

# SQLC
sqlc:
	@sqlc generate


# MIGRATIONS
up:
	@goose -dir ./repository/migrations postgres "user=user password=password dbname=db sslmode=disable" up

down:
	@goose -dir ./repository/migrations  postgres "user=user password=password dbname=db sslmode=disable" down

gooseinit:
	goose -dir ./repository/migrations create init sql