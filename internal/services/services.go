package services

import (
	"github.com/gin-gonic/gin"
	repository "github.com/srjchsv/simplebank/internal/repository/sqlc"
)

type Accounts interface {
	CreateAccount(*gin.Context)
	GetAccount(*gin.Context)
	UpdateAccount(*gin.Context)
	DeleteAccount(*gin.Context)
	ListAccounts(*gin.Context)
}

type Transfers interface {
	CreateTransfer(*gin.Context)
}

type Service struct {
	Accounts
	Transfers
}

func NewService(store repository.Store) *Service {
	return &Service{
		Accounts:  NewAccountsService(store),
		Transfers: NewTransfersService(store),
	}
}
