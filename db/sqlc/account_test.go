package db

import (
	"context"
	// "fmt"
	"testing"

	"github.com/dasotd/Ecom/util"
	_ "github.com/stretchr/testify"
	"github.com/stretchr/testify/require"
)



func createRandomAccount(t *testing.T) Account {
	user := createRandomUser(t)

	arg := CreateAccountParams{
		Owner:	user.Username,
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testBank.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	account := createRandomAccount(t)

	account2, err := testBank.GetAccount(context.Background(), account.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account2)
	require.Equal(t, account.ID, account2.ID)
	require.NotZero(t, account2.ID)

	// return account2
}

func TestDeleteAccount(t *testing.T){
	account := createRandomAccount(t)
	err := testBank.DeleteAccount(context.Background(), account.ID)
	require.NoError(t, err)

	accountz, err := testBank.GetAccount(context.Background(), account.ID)
	require.Error(t, err)
	require.Empty(t, accountz)

}

func TestListAccount(t *testing.T){
	var lastAccount Account
	for i := 0; i < 10; i++ {
		lastAccount = createRandomAccount(t)
	}

	args := ListAccountParams{
		Owner: lastAccount.Owner,
		Limit: 5,
		Offset: 0,
	}
	accounts, err := testBank.ListAccount(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, accounts)
	require.NotZero(t, accounts)
}

func TestUpdateAccount(t *testing.T){
	account := createRandomAccount(t)
	args := UpdateAccountParams{
		ID: account.ID,
		Balance: 500000,
	}

	account1, err := testBank.UpdateAccount(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, account1)
	require.Equal(t, account.ID, account1.ID)
}

func TestAddAccountBalance(t *testing.T){
	Account := createRandomAccount(t)
	args := AddAccountBalanceParams{
		Amount: 500090900090,
		ID: Account.ID,
	}

	account, err := testBank.AddAccountBalance(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, account.ID, account.ID)
}

func TestGetAccountForUpdate(t *testing.T) {
	account := createRandomAccount(t)

	account2, err := testBank.GetAccountForUpdate(context.Background(), account.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account2)
	require.Equal(t, account.ID, account2.ID)

	// return account2
}