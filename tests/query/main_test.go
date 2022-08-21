package repository

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	repository "github.com/srjchsv/simplebank/repository/sqlc"

	_ "github.com/lib/pq"
)

var (
	testQueries *repository.Queries
	testDB      *sql.DB
)

func TestMain(m *testing.M) {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("cannot load env file")
	}

	dbConfigs := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"))
	testDB, err = sql.Open(os.Getenv("DB_DRIVER"), dbConfigs)
	if err != nil {
		log.Fatal("cannot connect to db")
	}

	testQueries = repository.New(testDB)

	os.Exit(m.Run())
}
