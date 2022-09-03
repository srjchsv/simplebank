package services

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	repository "github.com/srjchsv/simplebank/internal/repository/sqlc"
	"github.com/srjchsv/simplebank/internal/services/responses"
)

type TransfersService struct {
	store repository.Store
}

func NewTransfersService(store repository.Store) *TransfersService {
	return &TransfersService{store: store}
}

type transferRequest struct {
	FromAccountID int64  `json:"from_account_id" binding:"required,min=1"`
	ToAccountID   int64  `json:"to_account_id" binding:"required,min=1"`
	Amount        int64  `json:"amount" binding:"required,gt=0"`
	Currency      string `json:"currency" binding:"required,currency"`
}

func (service *TransfersService) CreateTransfer(ctx *gin.Context) {
	var req transferRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.ErrorResponse(err))
		return
	}
	if !service.validAccount(ctx, req.FromAccountID, req.Currency) {
		return
	}
	if !service.validAccount(ctx, req.ToAccountID, req.Currency) {
		return
	}

	arg := repository.TransferTxParams{
		FromAccountID: req.FromAccountID,
		ToAccountID:   req.ToAccountID,
		Amount:        req.Amount,
	}
	result, err := service.store.TransferTx(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func (service *TransfersService) validAccount(ctx *gin.Context, accountID int64, currency string) bool {
	account, err := service.store.GetAccount(ctx, accountID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, responses.ErrorResponse(err))
			return false
		}
		ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse(err))
		return false
	}
	if account.Currency != currency {
		err := fmt.Errorf("account [%d] currency mismatch: %s vs %s", accountID, account.Currency, currency)
		ctx.JSON(http.StatusBadRequest, responses.ErrorResponse(err))
		return false
	}
	return true
}
