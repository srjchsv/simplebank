package server

import (
	"github.com/gin-gonic/gin"
	"github.com/srjchsv/simplebank/internal/services"
)

type Server struct {
	service *services.Service
	router  *gin.Engine
}

func NewServer(service *services.Service) *Server {
	server := &Server{service: service}
	router := gin.Default()

	router.POST("/accounts", service.Accounts.CreateAccount)
	router.GET("/accounts/:id", service.Accounts.GetAccount)
	router.DELETE("/accounts/:id",service.Accounts.DeleteAccount)
	router.GET("/accounts", service.Accounts.ListAccounts)
	

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
