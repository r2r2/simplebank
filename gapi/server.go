package gapi

import (
	"fmt"
	db "github.com/r2r2/simplebank/db/sqlc"
	"github.com/r2r2/simplebank/pb"
	"github.com/r2r2/simplebank/token"
	"github.com/r2r2/simplebank/util"
)

// Server serves gRPC requests
type Server struct {
	pb.UnimplementedSimpleBankServer
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
}

// NewServer creates a new gRPC server.
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	return server, nil
}
