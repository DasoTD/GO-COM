package db

import (
	"context"
	// "database/sql"
	"log"
	"os"
	"testing"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
)

var testBank Bank
// var testDB *sql.DB
const (
	DB_DRIVER="postgres"
	DB_SOURCE="postgresql://root:secret@localhost:5432/Ecom?sslmode=disable"
	
)


func TestMain(m *testing.M) {
	var err error
	connPool, err := pgxpool.New(context.Background(), DB_SOURCE)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testBank = NewBank(connPool)

	os.Exit(m.Run())
}