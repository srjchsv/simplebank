package main

import (
	"database/sql"

	"github.com/sirupsen/logrus"
	repository "github.com/srjchsv/simplebank/internal/repository/sqlc"
	"github.com/srjchsv/simplebank/internal/server"
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

	//initialize services and server
	store := repository.NewStore(conn)
	service := services.NewService(store)
	server := server.NewServer(service)

	//run server
	err = server.Start(config.ServersAddress)
	if err != nil {
		logrus.Fatal("cannot start server: ", err)
	}
}
