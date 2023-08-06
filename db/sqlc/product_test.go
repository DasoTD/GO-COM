package db

import (
	"context"
	"testing"

	_ "github.com/dasotd/Ecom/util"
	"github.com/stretchr/testify/require"
)

func TestCreateProduct(t *testing.T){
	args := CreateproductParams{
		Name: "SALAD ELEWE",
		Price: 20,
		Description: "Salad",

	}
	product, err := testBank.Createproduct(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, product)
}