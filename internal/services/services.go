package services

import (
	repository "github.com/srjchsv/simplebank/internal/repository/sqlc"
)

type Accounts interface {
	CreateAccount(req CreateAccountRequest) (repository.Account, error)
	GetAccount(req GetAccountRequest) (repository.Account, error)
	UpdateAccount(req UpdateAccountRequest) (repository.Account, error)
	DeleteAccount(req DeleteRequest) error
	ListAccounts(req ListAccountRequest) ([]repository.Account, error)
<<<<<<< HEAD
}

type Transfers interface {
	CreateTransfer(req TransferRequest) (repository.TransferTxResult, error)
=======
>>>>>>> 87ef88a (project structure)
}

type Transfers interface {
	CreateTransfer(req TransferRequest) (repository.TransferTxResult, error)
}

type Service struct {
	Accounts
	Transfers
}

func NewService(store repository.Store) *Service {
	return &Service{
		Accounts: NewAccountsService(store),
		Transfers: NewTransfersService(store),
	}
}
