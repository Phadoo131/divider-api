package db

import (
	"context"
	"strconv"
	"testing"

	"github.com/Phadoo131/divider-api/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomAccount(t *testing.T) Booker {
	arg := createBookerParam{
		Name: util.RandomName(),
		Email: util.RandomEmail(),
		PhoneNum: strconv.FormatInt(util.RandomInt(10000, 20000), 10),
		Password: util.RandomString(8),
		Address: util.RandomString(20),
	}

	account, err := testQuerie.CreateBookerAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.NotEmpty(t, account.Name)
	require.NotEmpty(t, account.Email)
	require.NotEmpty(t, account.PhoneNum)
	require.NotEmpty(t, account.Password)
	require.NotEmpty(t, account.Address)

	require.Equal(t, arg.Name, account.Name)
	require.Equal(t, arg.Email, account.Email)
	require.Equal(t, arg.PhoneNum, account.PhoneNum)
	require.Equal(t, arg.Password, account.Password)
	require.Equal(t, arg.Address, account.Address)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateBookerAccount (t *testing.T){
	CreateRandomAccount(t)
}