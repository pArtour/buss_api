package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/pArtour/buss_api/api"
	db "github.com/pArtour/buss_api/db/sqlc"
	"github.com/pArtour/buss_api/util"
	"log"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load configurations:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)
	err = server.Start(config.ServerAddress)

	if err != nil {
		log.Fatal("cannot start a server:", err)
	}
}
