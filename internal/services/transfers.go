package services

import (
	"context"
	"database/sql"
	"fmt"

	repository "github.com/srjchsv/simplebank/internal/repository/sqlc"
)

type TransferRequest struct {
	FromAccountID int64  `json:"from_account_id" binding:"required,min=1"`
	ToAccountID   int64  `json:"to_account_id" binding:"required,min=1"`
	Amount        int64  `json:"amount" binding:"required,gt=0"`
	Currency      string `json:"currency" binding:"required,currency"`
}

func (service *AccountsService) CreateTransfer(req TransferRequest) (repository.TransferTxResult, error) {
	fromAccount, err := service.validAccount(context.Background(), req.FromAccountID, req.Currency)
	if err != nil || !fromAccount {
		return repository.TransferTxResult{}, err
	}
	toAccount, err := service.validAccount(context.Background(), req.ToAccountID, req.Currency)
	if err != nil || !toAccount {
		return repository.TransferTxResult{}, err
	}
	arg := repository.TransferTxParams{
		FromAccountID: req.FromAccountID,
		ToAccountID:   req.ToAccountID,
		Amount:        req.Amount,
	}
	result, err := service.store.TransferTx(context.Background(), arg)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (service *AccountsService) validAccount(ctx context.Context, accountID int64, currency string) (bool, error) {
	account, err := service.store.GetAccount(ctx, accountID)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, sql.ErrNoRows
		}
		return false, err
	}
	if account.Currency != currency {
		err := fmt.Errorf("currency mismatch")
		return false, err
	}
	return true, nil
}
