package database

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {

	queries := &Queries{db: db}

	ctx := context.Background()

	arg := CreateAccountParams{
		Owner:       "Ilya",
		CountryCode: "RUs",
		Currency:    "RUB",
		Balance:     1000,
		Mail:        "ilya@hamil.com",
	}

	account, err := queries.CreateAccount(ctx, arg)
	require.NoErrorf(t, err, "Error with creating account %v", err)

	require.NotEmpty(t, account, "Returns empty struct")

}

func TestGetAccountByid(t *testing.T) {

	queries := &Queries{db: db}

	ctx := context.Background()

	id := int32(1)

	account, err := queries.GetAccounts(ctx, id)
	require.NotEmpty(t, account, "Returns empty struct")
	require.NoErrorf(t, err, "Error with get account %v", err)

}

func TestGetAccountsList(t *testing.T) {
	queries := &Queries{db: db}

	ctx := context.Background()

	limit := int32(1)

	accounts, err := queries.GetListAccounts(ctx, limit)
	require.NotEmpty(t, accounts, "Returns empty struct")
	require.NoErrorf(t, err, "Error with get accounts account %v", err)
}

func TestUpdateAcccount(t *testing.T) {

	queries := &Queries{db: db}

	ctx := context.Background()

	arg := UpdateAccountParams{
		ID:          1,
		Owner:       "Ilyass",
		CountryCode: "RUs",
		Currency:    "EUR",
		Balance:     2000,
		Mail:        "ilya@hamil.com",
	}

	account, err := queries.UpdateAccount(ctx, arg)
	require.NotEmpty(t, account, "Returns empty struct")
	require.NoErrorf(t, err, "Error with updating account %v", err)
}

func TestDeleteAccount(t *testing.T) {

	queries := &Queries{db: db}

	ctx := context.Background()

	id := int32(1)

	err := queries.DeleteAccount(ctx, id)

	require.NoError(t, err, "Error with delete account")

}





func TestCreateEntrie(t *testing.T) {

	queries := &Queries{db: db}

	ctx := context.Background()

	arg := CreateEntrieParams{
		AccountID: 1,
		Amount:    2531,
	}

	entry, err := queries.CreateEntrie(ctx, arg)
	require.NoErrorf(t, err, "Error with creating entrie %v", err)

	require.NotEmpty(t, entry, "Returns empty struct")

}



func TestUpdateEntrie(t *testing.T) {

	queries := &Queries{db: db}

	ctx := context.Background()

	arg := UpdateEntrieParams{
		ID: 1,
		AccountID: 1,
		Amount: 31,
	}

	entry, err := queries.UpdateEntrie(ctx, arg)
	require.NotEmpty(t, entry, "Returns empty struct")
	require.NoErrorf(t, err, "Error with updating entrie %v", err)
}


func TestGetEntrieByid(t *testing.T) {

	queries := &Queries{db: db}

	ctx := context.Background()

	id := int64(4)

	entry, err := queries.GetEntries(ctx, id)
	require.NotEmpty(t, entry, "Returns empty struct")
	require.NoErrorf(t, err, "Error with get entrie %v", err)

}

func TestGetEntriesList(t *testing.T) {
	queries := &Queries{db: db}

	ctx := context.Background()

	limit := int32(1)

	accounts, err := queries.GetListEntries(ctx, limit)
	require.NotEmpty(t, accounts, "Returns empty struct")
	require.NoErrorf(t, err, "Error with get entries list %v", err)
}

func TestDeleteEntrie(t *testing.T) {

	queries := &Queries{db: db}

	ctx := context.Background()

	id := int64(5)

	err := queries.DeleteEntrie(ctx, id)

	require.NoError(t, err, "Error with delete entrie")

}





func TestCreateTransfer(t *testing.T){

	queries :=&Queries{db: db}

	ctx := context.Background()

	arg := CreateTransferParams{
		FromAccountID: 1,
		ToAccountID: 2,
		Amount: 200,
	}
	transfer, err := queries.CreateTransfer(ctx, arg)

	require.NotEmpty(t, transfer, "Returns empty struct")
	require.NoErrorf(t, err, "Error with create transfer  %v", err)

}


func TestUpdateTransfer(t *testing.T){
	queries := &Queries{db: db}

	ctx := context.Background()


	arg := UpdateTransferParams{
		ID: 1,
		FromAccountID: 1,
		ToAccountID: 2,
		Amount: 300,
	}
	transfer, err := queries.UpdateTransfer(ctx, arg)

	require.NotEmpty(t, transfer, "Returns empty struct")
	require.NoErrorf(t, err, "Error with update transfer  %v", err)
}


func TestGetTransferById(t *testing.T){

	queries := &Queries{db: db}

	ctx := context.Background()

	id := int64(1)

	transfer, err := queries.GetTransfers(ctx, id)


	require.NotEmpty(t, transfer, "Returns empty struct")
	require.NoErrorf(t, err, "Error with get transfer by id  %v", err)

}


func TestGetListTransfers(t *testing.T){
	queries := &Queries{db:db}


	ctx := context.Background()


	limit := int32(1)
	transfer, err := queries.GetListTransfers(ctx, limit)


	require.NotEmpty(t, transfer, "Returns empty struct")
	require.NoErrorf(t, err, "Error with get transfer list  %v", err)

}

func TestDeleteTransfer(t *testing.T){
	queries := &Queries{db:db}


	ctx := context.Background()


	id := int64(1)
	err := queries.DeleteTransfer(ctx, id)
	require.NoErrorf(t, err, "Error with delete transfer  %v", err)
}