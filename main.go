package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/tuanbui-n9/simplebank/api"
	db "github.com/tuanbui-n9/simplebank/db/sqlc"
	"github.com/tuanbui-n9/simplebank/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatalf("cannot load config: %v", err)
	}
	conn, err := pgxpool.New(context.Background(), config.DBSource)

	if err != nil {
		log.Fatalf("cannot connect to db: %v", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatalf("cannot create server: %v", err)
	}

	err = server.Start(config.ServerAddress)

	if err != nil {
		log.Fatalf("cannot start server: %v", err)
	}
}
