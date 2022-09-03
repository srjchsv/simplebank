package main

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/srjchsv/simplebank/internal/handler"
	repository "github.com/srjchsv/simplebank/internal/repository/sqlc"
	"github.com/srjchsv/simplebank/internal/services"
	"github.com/srjchsv/simplebank/util"

	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		logrus.Fatal("cannot load config: ", err)
	}
	//db connection
	conn, err := sql.Open(config.DbDriver, config.PgUrl)
	if err != nil {
		logrus.Fatal("cannot connect to db", err)
	}
	// Set db connection pool
	conn.SetMaxOpenConns(config.PgPool)
	conn.SetMaxIdleConns(config.PgPool)

	//initialize store, services and server
	store := repository.NewStore(conn)
	services := services.NewService(store)
	handlers := handler.NewHandler(services)

	//run server
	r := gin.Default()

	
	handlers.InitRouter(r)
	logrus.Fatal(r.Run(config.ServersAddress))
}
