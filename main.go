package main

import (
	"context"
	"log"
	"net"

	"github.com/dasotd/Ecom/api"
	db "github.com/dasotd/Ecom/db/sqlc"
	"github.com/dasotd/Ecom/gapi"
	"github.com/dasotd/Ecom/pb"
	"github.com/dasotd/Ecom/util"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)



func main(){
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("can not load config")
	}
	connPool, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	bank := db.NewBank(connPool)
	runGrpcServer(config, bank)
}

func runGrpcServer(config util.Config, bank db.Bank) {
	server, err := gapi.NewServer(config, bank)
	if err != nil {
		log.Fatal("cannot create server", err) //.Err(err).Msg("cannot create server")
	}

	gprcLogger := grpc.UnaryInterceptor(gapi.GrpcLogger)
	grpcServer := grpc.NewServer(gprcLogger)
	pb.RegisterEcomServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatal("cannot create server", err ) //.Err(err).Msg("cannot create listener")
	}

	log.Printf("start gRPC server at %s", listener.Addr().String()) //.Msgf("start gRPC server at %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start Grpc server", err) //.Err(err).Msg("cannot start gRPC server")
	}
}

func runGinServer(config util.Config, bank db.Bank) {
	server, err := api.NewServer(config, bank)
	if err != nil {
		log.Fatal("cannot create server", err)//.Err(err).Msg("cannot create server")
	}

	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal("cannot start server", err) //.Err(err).Msg("cannot start server")
	}
}