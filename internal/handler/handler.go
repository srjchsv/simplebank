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
	auth := app.Group("/auth")
	auth.POST("/signup", h.CreateAccount)
	auth.POST("/signin", h.SignIn)

	accounts := app.Group("/accounts", h.UserIdentity)
	accounts.GET("/:id", h.GetAccount)
	accounts.PUT("/:id", h.UpdateAccount)
	accounts.DELETE("/:id", h.DeleteAccount)
	accounts.GET("", h.ListAccounts)
	accounts.POST("/transfers", h.CreateTransfer)

	return app
}
