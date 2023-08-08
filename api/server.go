package api

import (
	"github.com/dasotd/Ecom/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)


type Server struct {
	bank db.Bank
	router *gin.Engine
}

func NewServer( bank db.Bank) *Server {
	server := &Server{bank: bank}
	router := gin.Default()


	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	router.POST("/account", server.createAccount)
	router.GET("account/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)
	router.POST("/transfers", server.createTransfer)
	router.POST("user", server.createUser)
	

	server.router = router
	return server
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}


// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error{
	return server.router.Run(address)
}
