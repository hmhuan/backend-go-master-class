package main

import (
	"database/sql"
	"log"

	"github.com/hmhuan/backend-go-master-class/api"
	db "github.com/hmhuan/backend-go-master-class/db/sqlc"
	"github.com/hmhuan/backend-go-master-class/util"
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
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("failed to start server: ", err)
	}

}
