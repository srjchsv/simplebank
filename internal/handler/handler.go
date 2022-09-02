package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/srjchsv/simplebank/internal/services"
)

type Handler struct {
	services *services.Service
}

func NewHandler(services *services.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRouter(app *gin.Engine) *gin.Engine {
	app.POST("/accounts", h.services.Accounts.CreateAccount)
	app.GET("/accounts/:id", h.services.Accounts.GetAccount)
	app.DELETE("/accounts/:id", h.services.Accounts.DeleteAccount)
	app.GET("/accounts", h.services.Accounts.ListAccounts)

	return app
}
