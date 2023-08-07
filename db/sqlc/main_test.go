package db

import (
	"context"
	// "database/sql"
	"log"
	"os"
	"testing"

	"github.com/dasotd/Ecom/util"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
)

var testBank Bank



func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load environment variable", err)
	}
	connPool, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testBank = NewBank(connPool)

	os.Exit(m.Run())
}