package services

import (
	"context"

	repository "github.com/srjchsv/simplebank/internal/repository/sqlc"
)

type AccountsService struct {
	store repository.Store
}

func NewAccountsService(store repository.Store) *AccountsService {
	return &AccountsService{store: store}
}

type CreateAccountRequest struct {
	Owner    string `json:"owner" binding:"required"`
	Balance  int64  `json:"balance" binding:"required,min=0"`
	Currency string `json:"currency" binding:"required,currency"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (service *AccountsService) CreateAccount(req CreateAccountRequest) (repository.Account, error) {
	arg := repository.CreateAccountParams{
		Owner:    req.Owner,
		Balance:  req.Balance,
		Currency: req.Currency,
	}
	account, err := service.store.CreateAccount(context.Background(), arg)
	if err != nil {
		return account, err
	}
	return account, nil
}

type GetAccountRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (service *AccountsService) GetAccount(req GetAccountRequest) (repository.Account, error) {
	account, err := service.store.GetAccount(context.Background(), req.ID)
	if err != nil {
		return account, err
	}
	return account, nil
}

type UpdateAccountRequest struct {
	Owner   string `json:"owner" binding:"required"`
	Balance int64  `json:"balance" binding:"required,min=0"`
	ID      int64  `uri:"id" binding:"required,min=1"`
}

func (service *AccountsService) UpdateAccount(req UpdateAccountRequest) (repository.Account, error) {
	arg := repository.UpdateAccountParams{
		Owner:   req.Owner,
		Balance: req.Balance,
		ID:      req.ID,
	}
	account, err := service.store.UpdateAccount(context.Background(), arg)
	if err != nil {
		return account, err
	}
	return account, nil
}

type DeleteRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (service *AccountsService) DeleteAccount(req DeleteRequest) error {
	err := service.store.DeleteAccount(context.Background(), req.ID)
	if err != nil {
		return err
	}
	return nil
}

type ListAccountRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (service *AccountsService) ListAccounts(req ListAccountRequest) ([]repository.Account, error) {
	arg := repository.ListAccountsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}
	account, err := service.store.ListAccounts(context.Background(), arg)
	if err != nil {
		return account, err
	}
	return account, nil
}
