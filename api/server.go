package api

import (
	"fmt"

	"github.com/dasotd/Ecom/db/sqlc"
	"github.com/dasotd/Ecom/token"
	"github.com/dasotd/Ecom/util"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)


type Server struct {
	config util.Config
	bank db.Bank
	tokenMaker token.Maker
	router *gin.Engine
}

func NewServer(config util.Config, bank db.Bank) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("Cannot create token %w", err)

	}
	server := &Server{
		config: config,
		bank: bank,
		tokenMaker: tokenMaker,
	}


	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	
	

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter(){
	router := gin.Default()
	router.POST("/account", server.createAccount)
	router.GET("account/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)
	router.POST("/transfers", server.createTransfer)
	router.POST("/user", server.createUser)
	router.POST("/user/login", server.loginUser)

}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}


// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error{
	return server.router.Run(address)
}
