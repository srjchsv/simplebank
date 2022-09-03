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
	accounts.POST("", h.services.Accounts.CreateAccount)
	accounts.GET("/:id", h.services.Accounts.GetAccount)
	accounts.PUT("/:id", h.services.Accounts.UpdateAccount)
	accounts.DELETE("/:id", h.services.Accounts.DeleteAccount)
	accounts.GET("", h.services.Accounts.ListAccounts)

	//Transfers route
	transfers := app.Group("/transfers")
	transfers.POST("", h.services.Transfers.CreateTransfer)

	return app
}
