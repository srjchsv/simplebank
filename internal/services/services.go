package services

import (
	"github.com/gin-gonic/gin"
	repository "github.com/srjchsv/simplebank/internal/repository/sqlc"
)

type Accounts interface {
	CreateAccount(*gin.Context)
	GetAccount(*gin.Context)
	DeleteAccount(*gin.Context)
	ListAccounts(*gin.Context)
}

type Service struct {
	Accounts
}

func NewService(store repository.Store) *Service {
	return &Service{Accounts: NewAccountsService(store)}
}
