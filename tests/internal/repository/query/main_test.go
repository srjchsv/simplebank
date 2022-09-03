package repository

import (
	"database/sql"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	repository "github.com/srjchsv/simplebank/internal/repository/sqlc"
	"github.com/srjchsv/simplebank/util"

	_ "github.com/lib/pq"
)

var (
	testQueries *repository.Queries
	testDB      *sql.DB
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	config, err := util.LoadConfig("../../../..")
	if err != nil {
		logrus.Fatal("cannot load config: ", err)
	}
	testDB, err = sql.Open(config.DbDriver, config.PgUrl)
	if err != nil {
		logrus.Fatal("cannot connect to db", err)
	}

	testQueries = repository.New(testDB)

	os.Exit(m.Run())
}
