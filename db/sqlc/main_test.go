package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

var testQueries *Queries
var testDb *pgxpool.Pool

const (
	dbSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
)

func TestMain(m *testing.M) {
	var err error
	testDb, err = pgxpool.New(context.Background(), dbSource)

	if err != nil {
		log.Fatalf("cannot connect to db: %v", err)
	}

	testQueries = New(testDb)

	os.Exit(m.Run())
}
