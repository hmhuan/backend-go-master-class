package main

import (
	"database/sql"
	"log"

	"github.com/hmhuan/simple-bank/api"
	db "github.com/hmhuan/simple-bank/db/sqlc"
	"github.com/hmhuan/simple-bank/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Failed to load configuration", err)
	}

	dbConn, err := sql.Open(config.DbDriver, config.DbSource)
	if err != nil {
		log.Fatal("failed to connect to database: ", err)
	}

	store := db.NewStore(dbConn)
	server, err := api.NewServer(*config, store)
	if err != nil {
		log.Fatal("failed to create server: ", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("failed to start server: ", err)
	}

}
