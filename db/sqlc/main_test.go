package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/leijunjiang/simple_bank/util"
)

var testQueries *Queries
var testDB *pgxpool.Pool

func TestMain(m *testing.M) {
	var err error
	config, configErr := util.LoadConfig("../..")
	if configErr != nil {
		log.Fatal("cannot load config:", configErr)
	}

	testDB, err = pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	testQueries = New(testDB)

	code := m.Run()

	testDB.Close()

	os.Exit(code)
}
