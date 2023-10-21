package db

import (
	"context"
	"log"
	"os"
	"testing"

	_"github.com/Phadoo131/divider-api/util"
	_"github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	_"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)
var testQuerie *Queries
var dbSource = os.Getenv("DB_SOURCE")

func TestMain(m *testing.M){
	var err error

	// config, err := util.LoadConfig("../..")
	// if err != nil {
	// 	log.Fatal("cannot load config:", err)
	// }

	conn, err := pgxpool.New(context.Background(), dbSource)
	if err != nil {
		log.Fatal("Unable to connect to database: ", err)
		os.Exit(1)
	}
	testQuerie = New(conn)

	os.Exit(m.Run())
}