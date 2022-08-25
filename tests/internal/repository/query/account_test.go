package repository

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/pioz/faker"
	repository "github.com/srjchsv/simplebank/internal/repository/sqlc"
	"github.com/srjchsv/simplebank/util"
	"github.com/stretchr/testify/require"
)

func TestRepository_CreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestRepository_UpdateAccount(t *testing.T) {
	
	account1 := createRandomAccount(t)

	arg := repository.UpdateAccountParams{
		ID:      account1.ID,
		Owner:   account1.Owner,
		Balance: faker.Int64InRange(555, 888),
	}
	updated, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, updated)
	require.Equal(t, account1.ID, updated.ID)
	require.Equal(t, account1.Owner, updated.Owner)
	require.Equal(t, arg.Balance, updated.Balance)
	require.WithinDuration(t, account1.CreatedAt, updated.CreatedAt, time.Second)
}
func TestRepository_GetAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	account2, err := testQueries.GetAccount(context.Background(), account1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, account2)
	require.Equal(t, account1.Balance, account2.Balance)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, account1.Currency, account2.Currency)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}

func TestQuery_DeleteAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), account1.ID)
	require.NoError(t, err)

	deleted, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, deleted)
}

func TestQuery_ListAccount(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}
	arg := repository.ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}
	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 5)
	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}

func createRandomAccount(t *testing.T) repository.Account {
	arg := repository.CreateAccountParams{
		Owner:    faker.FirstName(),
		Balance:  faker.Int64InRange(100, 1000),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)
	require.Equal(t, arg.Owner, account.Owner)
	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}
