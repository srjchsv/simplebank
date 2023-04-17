package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/srjchsv/simplebank/internal/handler"
	repository "github.com/srjchsv/simplebank/internal/repository/sqlc"
	"github.com/srjchsv/simplebank/internal/services"
	"github.com/srjchsv/simplebank/util"
)

type DbConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DbName   string
	Pool     string
}

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		logrus.Fatal("cannot load config: ", err)
	}

	dbConfig := DbConfig{
		Host:     config.PgHost,
		Username: config.PgUsername,
		Password: config.PgPassword,
		DbName:   config.PgName,
		Port:     config.PgPort,
	}
	dbUrl := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.DbName)
	//db connection
	conn, err := sqlx.Open(config.DbDriver, dbUrl)
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
