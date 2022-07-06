package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/hmhuan/backend-go-master-class/util"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Failed to load configuration", err)
	}

	testDB, err = sql.Open(config.DbDriver, config.DbSource)
	if err != nil {
		log.Fatal("failed to connect to database: ", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
