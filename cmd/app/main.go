package main

import (
	"database/sql"
	"fmt"

	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/srjchsv/simplebank/internal/server"
	"github.com/srjchsv/simplebank/internal/services"
	repository "github.com/srjchsv/simplebank/repository/sqlc"

	_ "github.com/lib/pq"
)

func main() {
	//load environment variables
	err := godotenv.Load(".env")
	if err != nil {
		logrus.Fatal("cannot load env file")
	}

	//db configs
	dbConfigs := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"))
	//db connection
	conn, err := sql.Open(os.Getenv("DB_DRIVER"), dbConfigs)
	if err != nil {
		logrus.Fatal("cannot connect to db")
	}
	// Set db connection pool
	pool, err := strconv.Atoi(os.Getenv("POSTGRES_POOL"))
	if err != nil {
		logrus.Fatal("error getting db pool env variable: ", err)
	}
	conn.SetMaxOpenConns(pool)
	conn.SetMaxIdleConns(pool)

	//initialize services and server
	store := repository.NewStore(conn)
	service := services.NewService(store)
	server := server.NewServer(service)

	//run server
	err = server.Start(os.Getenv("ADDRESS"))
	if err != nil {
		logrus.Fatal("cannot start server: ", err)
	}
}
