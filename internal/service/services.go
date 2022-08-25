package services

import repository "github.com/srjchsv/simplebank/internal/repository/sqlc"

type Service struct {
	store *repository.Store
}

func NewService(store *repository.Store) *Service {
	return &Service{store: store}
}
