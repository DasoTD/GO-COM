package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateCart(t *testing.T){
	user := createRandomUser(t)
	args := CreateCartParams{
		Owner: user.Username,
		Product: "Pasta",
		Quantity: 12,
	}

	cart, err := testBank.CreateCart(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, cart)

}