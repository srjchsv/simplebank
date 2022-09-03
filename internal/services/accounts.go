package services

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	repository "github.com/srjchsv/simplebank/internal/repository/sqlc"
	"github.com/srjchsv/simplebank/internal/services/responses"
)

type AccountsService struct {
	store repository.Store
}

func NewAccountsService(store repository.Store) *AccountsService {
	return &AccountsService{store: store}
}

type createAccountRequest struct {
	Owner    string `json:"owner" binding:"required"`
	Balance  int64  `json:"balance" binding:"required,min=0"`
	Currency string `json:"currency" binding:"required,currency"`
}

func (service *AccountsService) CreateAccount(ctx *gin.Context) {
	var req createAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.ErrorResponse(err))
		return
	}
	arg := repository.CreateAccountParams{
		Owner:    req.Owner,
		Balance:  req.Balance,
		Currency: req.Currency,
	}
	account, err := service.store.CreateAccount(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, account)
}

type getAccountRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (service *AccountsService) GetAccount(ctx *gin.Context) {
	var req getAccountRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.ErrorResponse(err))
		return
	}
	account, err := service.store.GetAccount(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, responses.ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, account)
}

type updateAccountRequest struct {
	Owner   string `json:"owner" binding:"required"`
	Balance int64  `json:"balance" binding:"required,min=0"`
	ID      int64  `uri:"id" binding:"required,min=1"`
}

func (service *AccountsService) UpdateAccount(ctx *gin.Context) {
	var req updateAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.ErrorResponse(err))
		return
	}
	arg := repository.UpdateAccountParams{
		Owner:   req.Owner,
		Balance: req.Balance,
		ID:      req.ID,
	}

	account, err := service.store.UpdateAccount(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, account)
}

type deleteRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (service *AccountsService) DeleteAccount(ctx *gin.Context) {
	var req deleteRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.ErrorResponse(err))
		return
	}
	err := service.store.DeleteAccount(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, responses.ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("account %v deleted successfully", req.ID),
	})
}

type listAccountRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (service *AccountsService) ListAccounts(ctx *gin.Context) {
	var req listAccountRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.ErrorResponse(err))
		return
	}
	arg := repository.ListAccountsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}
	accounts, err := service.store.ListAccounts(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, accounts)
}
