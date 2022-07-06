package main

import (
	"database/sql"
	"log"

	"github.com/hmhuan/backend-go-master-class/api"
	db "github.com/hmhuan/backend-go-master-class/db/sqlc"
	_ "github.com/lib/pq"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgres://root:secret@localhost:5432/simple_bank?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	dbConn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("failed to connect to database: ", err)
	}

	store := db.NewStore(dbConn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("failed to start server: ", err)
	}

}
