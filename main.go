package main

import (
	"context"
	"log"

	"github.com/dasotd/Ecom/api"
	db "github.com/dasotd/Ecom/db/sqlc"
	"github.com/dasotd/Ecom/util"
	"github.com/jackc/pgx/v5/pgxpool"
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
	server :=api.NewServer(bank)
	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal("cannot start sever:", err)
	}
}