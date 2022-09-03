package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/srjchsv/simplebank/internal/services"
	"github.com/srjchsv/simplebank/internal/services/validate"
)

type Handler struct {
	services *services.Service
}

func NewHandler(services *services.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRouter(app *gin.Engine) *gin.Engine {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validate.ValidCurrency)
	}
	//Accounts route
	accounts := app.Group("/accounts")
	accounts.POST("", h.CreateAccount)
	accounts.GET("/:id", h.GetAccount)
	accounts.PUT("/:id", h.UpdateAccount)
	accounts.DELETE("/:id", h.DeleteAccount)
	accounts.GET("", h.ListAccounts)
	//Transfers route
	transfers := app.Group("/transfers")
	transfers.POST("", h.CreateTransfer)

	return app
}
