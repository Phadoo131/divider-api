package db

import (
	"testing"
	"log"
	"os"
	"database/sql"
	"context"

	_"github.com/joho/godotenv"
	_"github.com/lib/pq"
	"github.com/jackc/pgx/v5"
	_"github.com/jackc/pgx/v5/pgconn"
)
var testQuerie *Queries
var testDB *sql.DB

var dbSource = os.Getenv("dbSource")

func TestMain(m *testing.M){
	var err error

	testDB, err := pgx.Connect(context.Background(), dbSource)
	if err != nil {
		log.Fatal("Unable to connect to database: ", err)
		os.Exit(1)
	}
	defer testDB.Close(context.Background())

	testQuerie = New(testDB)

	os.Exit(m.Run())
}