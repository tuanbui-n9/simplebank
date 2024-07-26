package gapi

import (
	"fmt"

	db "github.com/tuanbui-n9/simplebank/db/sqlc"
	"github.com/tuanbui-n9/simplebank/pb"
	"github.com/tuanbui-n9/simplebank/token"
	"github.com/tuanbui-n9/simplebank/util"
	"github.com/tuanbui-n9/simplebank/worker"
)

// Server serves gRPC requests for our bank service
type Server struct {
	pb.UnimplementedSimpleBankServer
	config          util.Config
	store           db.Store
	tokenMaker      token.Maker
	taskDistributor worker.TaskDistributor
}

// NewServer creates a new server
func NewServer(config util.Config, store db.Store, taskDistributor worker.TaskDistributor) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:          config,
		store:           store,
		tokenMaker:      tokenMaker,
		taskDistributor: taskDistributor,
	}

	return server, nil
}
