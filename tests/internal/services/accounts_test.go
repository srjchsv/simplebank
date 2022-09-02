package services

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/pioz/faker"
	"github.com/srjchsv/simplebank/internal/handler"
	repository "github.com/srjchsv/simplebank/internal/repository/sqlc"

	"github.com/srjchsv/simplebank/internal/services"
	repoMock "github.com/srjchsv/simplebank/tests/internal/repository/sqlc/mock"
	"github.com/srjchsv/simplebank/util"
	"github.com/stretchr/testify/require"
)

func TestServices_GetAccount(t *testing.T) {
	account := randomAccount()

	tests := []struct {
		name          string
		accountID     int64
		buildStubs    func(store *repoMock.MockStore)
		checkResponse func(t *testing.T, w *httptest.ResponseRecorder)
	}{
		{
			name:      "OK",
			accountID: account.ID,
			buildStubs: func(store *repoMock.MockStore) {
				store.EXPECT().
					GetAccount(gomock.Any(), gomock.Eq(account.ID)).
					Times(1).
					Return(account, nil)
			},
			checkResponse: func(t *testing.T, w *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, w.Code)
				requireBodyMatchAccount(t, w.Body, account)
			},
		},
		{
			name:      "NotFound",
			accountID: account.ID,
			buildStubs: func(store *repoMock.MockStore) {
				store.EXPECT().
					GetAccount(gomock.Any(), gomock.Eq(account.ID)).
					Times(1).
					Return(repository.Account{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, w *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, w.Code)
			},
		},
		{
			name:      "InternalError",
			accountID: account.ID,
			buildStubs: func(store *repoMock.MockStore) {
				store.EXPECT().
					GetAccount(gomock.Any(), gomock.Eq(account.ID)).
					Times(1).
					Return(repository.Account{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, w *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, w.Code)
			},
		},
		{
			name:      "InvalidID",
			accountID: 0,
			buildStubs: func(store *repoMock.MockStore) {
				store.EXPECT().
					GetAccount(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, w *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, w.Code)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			store := repoMock.NewMockStore(c)
			test.buildStubs(store)

			services := services.NewService(store)
			handler := handler.NewHandler(services)

			r := gin.New()
			handler.InitRouter(r)

			w := httptest.NewRecorder()
			url := fmt.Sprintf("/accounts/%d", test.accountID)
			req, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			r.ServeHTTP(w, req)
			test.checkResponse(t, w)
		})
	}
}

func requireBodyMatchAccount(t *testing.T, body *bytes.Buffer, account repository.Account) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)

	var gotAccount repository.Account
	err = json.Unmarshal(data, &gotAccount)
	require.NoError(t, err)
	require.Equal(t, account, gotAccount)

}

func randomAccount() repository.Account {
	return repository.Account{
		ID:       util.RandomInt(1, 1000),
		Owner:    faker.FirstName(),
		Balance:  util.RandomInt(333, 777),
		Currency: util.RandomCurrency(),
	}
}
