package main

import (
	"context"
	"log"
	"net"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/tuanbui-n9/simplebank/api"
	db "github.com/tuanbui-n9/simplebank/db/sqlc"
	"github.com/tuanbui-n9/simplebank/gapi"
	"github.com/tuanbui-n9/simplebank/pb"
	"github.com/tuanbui-n9/simplebank/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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
	runGrpcServer(config, store)
}

func runGrpcServer(config util.Config, store db.Store) {
	server, err := gapi.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterSimpleBankServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.GrpcServerAddress)
	if err != nil {
		log.Fatalf("cannot listen to address: %v", err)
	}

	log.Printf("start grpc server on %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start grpc server")
	}
}

func runGinServer(config util.Config, store db.Store) {
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatalf("cannot create server: %v", err)
	}

	err = server.Start(config.HttpServerAddress)

	if err != nil {
		log.Fatalf("cannot start server: %v", err)
	}
}
