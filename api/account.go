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
	if err := ctx.BindJSON(&req); err != nil {
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

type getAccountRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getAccount(ctx *gin.Context){

	var req getAccountRequest
	if err := ctx.BindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	account, err := server.bank.GetAccount(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}
	ctx.JSON(http.StatusOK, account)

}