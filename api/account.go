package api

import (
	"context"
	"net/http"

	db "github.com/dasotd/Ecom/db/sqlc"
	// "github.com/dasotd/Ecom/util"
	"github.com/gin-gonic/gin"
)
type createAccountRequest struct {
	Owner string `json:"owner" binding:"required"`
	Currency string `json:"currency" binding:"required,currency"` //oneof=GBP EUR USD CAD
}

func (server *Server) createAccount( ctx *gin.Context){
	var req createAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateAccountParams{
		Owner: req.Owner,
		Balance: 5000000,
		Currency: req.Currency,
	}

	account, err:= server.bank.CreateAccount(context.Background(), arg)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}
	ctx.JSON(http.StatusCreated, account)
}