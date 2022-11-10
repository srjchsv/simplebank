package handler

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/srjchsv/simplebank/internal/services"
	"github.com/srjchsv/simplebank/internal/services/responses"
)

func (h *Handler) CreateAccount(ctx *gin.Context) {
	var req services.CreateAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.ErrorResponse(err))
		return
	}

	id, err := SignUp(req.Owner, req.Username, req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse(err))
		return
	}

	account, err := h.services.Accounts.CreateAccount(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"account": account,
		"id":      id,
	})
}

func (h *Handler) GetAccount(ctx *gin.Context) {
	var req services.GetAccountRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.ErrorResponse(err))
		return
	}
	account, err := h.services.Accounts.GetAccount(req)
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

func (h *Handler) UpdateAccount(ctx *gin.Context) {
	var req services.UpdateAccountRequest
	userID := ctx.GetInt("UserID")
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.ErrorResponse(err))
		return
	}
	logrus.Info(req.ID, userID)
	if int(req.ID) != userID {
		ctx.JSON(http.StatusBadRequest, responses.ErrorResponse(errors.New("wrong id")))
		return
	}
	account, err := h.services.Accounts.UpdateAccount(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, account)
}

func (h *Handler) DeleteAccount(ctx *gin.Context) {
	var req services.DeleteRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.ErrorResponse(err))
		return
	}
	userID := ctx.GetInt("UserID")
	if int(req.ID) != userID {
		ctx.JSON(http.StatusBadRequest, responses.ErrorResponse(errors.New("wrong id")))
		return
	}
	err := h.services.Accounts.DeleteAccount(req)
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

func (h *Handler) ListAccounts(ctx *gin.Context) {
	var req services.ListAccountRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.ErrorResponse(err))
		return
	}
	accounts, err := h.services.Accounts.ListAccounts(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, accounts)
}
