package si_bank

import (
	"database/sql"
	"log"
	"os"
	"testing"
	_ "github.com/lib/pq"
)

const (
	driverName = "postgres"
	dataSourceName = "postgresql://postgres:postgres@localhost:5432/si_bank?sslmode=disable"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	conn, err := sql.Open(driverName, dataSourceName)

	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries =  New(conn)
	os.Exit(m.Run())
}
