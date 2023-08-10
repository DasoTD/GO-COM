package gapi

import (
	"fmt"

	db "github.com/dasotd/Ecom/db/sqlc"
	"github.com/dasotd/Ecom/pb"
	"github.com/dasotd/Ecom/token"
	"github.com/dasotd/Ecom/util"
	// "github.com/dasotd/Ecom/worker"
)

// Server serves gRPC requests for our banking service.
type Server struct {
	pb.UnimplementedEcomServer
	config     util.Config
	bank      db.Bank
	tokenMaker token.Maker
	// taskDistributor worker.TaskDistributor
}

// mustEmbedUnimplementedEcomServer implements pb.EcomServer.
func (*Server) mustEmbedUnimplementedEcomServer() {
	panic("unimplemented")
}

// NewServer creates a new gRPC server.
func NewServer(config util.Config, bank db.Bank) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		bank:      bank,
		tokenMaker: tokenMaker,
		// taskDistributor: taskDistributor,
	}

	return server, nil
}
