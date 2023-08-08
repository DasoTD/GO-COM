package api 

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	mockdb "github.com/dasotd/Ecom/db/mock"
	db "github.com/dasotd/Ecom/db/sqlc"
	"github.com/dasotd/Ecom/util"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func randomAccount(owner string) db.Account{
	return db.Account{
		ID: util.RandomInt(1,1000),
		Owner: owner,
		Balance: util.RandomMoney(),
		Currency: util.RandomCurrency(),
		CreatedAt: time.Now(),
	}
}

func TestGetAccount(t *testing.T) {
	user, _ := randomUser(t)
	account := randomAccount(user.Username)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	bank := mockdb.NewMockBank(ctrl)
	bank.EXPECT().
		GetAccount(gomock.Any(), gomock.Eq(account.ID)).
		Times(1).
		Return(account, nil).AnyTimes()
	
	server := newTestServer(t,bank)
	recorder := httptest.NewRecorder()

	url := fmt.Sprintf("/accounts/%d", account.ID)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	require.NoError(t, err)
	server.router.ServeHTTP(recorder, request)
	fmt.Print("v", recorder.Code)
	// require.Equal(t, http.StatusOK, recorder.Code)
	// requireBodyMatchAccounts(t, recorder.Body, account)
}


func requireBodyMatchAccounts(t *testing.T, body *bytes.Buffer, accounts db.Account) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotAccounts []db.Account
	err = json.Unmarshal(data, &gotAccounts)
	require.NoError(t, err)
	require.Equal(t, accounts, gotAccounts)
}
