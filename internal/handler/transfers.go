package handler

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/srjchsv/simplebank/internal/services"
	"github.com/srjchsv/simplebank/internal/services/responses"
)

func (h *Handler) CreateTransfer(ctx *gin.Context) {
	var req services.TransferRequest
	userID := ctx.GetInt("UserID")
	if int(req.FromAccountID) != userID {
		ctx.JSON(http.StatusBadRequest, responses.ErrorResponse(errors.New("wrong id")))
		return
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.ErrorResponse(err))
		return
	}
	result, err := h.services.Accounts.CreateTransfer(req)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, responses.ErrorResponse(err))
			return
		}
		if err == fmt.Errorf("currency mismatch") {
			ctx.JSON(http.StatusBadRequest, responses.ErrorResponse(err))
		}
		ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, result)
}
