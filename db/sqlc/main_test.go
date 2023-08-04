package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	// "github.com/dasotd/simplebank/util"
	// "github.com/dasotd/ecom/util"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB
const (
	DB_DRIVER="postgres"
	DB_SOURCE="postgresql://root:secret@localhost:5432/Ecom?sslmode=disable"
	
)


func TestMain(m *testing.M) {
	// config, err := util.LoadConfig("../..")
	// if err != nil {
	// 	log.Fatal("cannot load config:", err)
	// }
	// DB_DRIVER=postgres
	// DB_SOURCE=postgresql://root:secret@localhost:5432/test?sslmode=disable
	
	testDB, err := sql.Open(DB_DRIVER, DB_SOURCE)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}