package db

import (
	"context"
	"database/sql"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func CreateRandomEntry(t *testing.T) Entry {
	account := CreateRandomAccount(t)
	arg := CreateEntrieParams{
		AccountID: account.ID,
		Amount:    sql.NullInt64{},
	}

	entry, err := testQueries.CreateEntrie(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, arg.Amount, entry.Amount)
	require.Equal(t, arg.AccountID, entry.AccountID)

	require.NotZero(t, entry.AccountID)
	require.NotZero(t, entry.CreatedAt)

	return entry
}

func TestQueries_CreateEntrie(t *testing.T) {
	CreateRandomEntry(t)
}

func TestQueries_GetEntrie(t *testing.T) {
	entry1 := CreateRandomEntry(t)
	entry2, err := testQueries.GetEntrie(context.Background(), entry1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.Amount, entry2.Amount)
	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.WithinDuration(t, entry1.CreatedAt, entry2.CreatedAt, time.Second)
}

func TestQueries_UpdateEntrie(t *testing.T) {
	entry1 := CreateRandomEntry(t)

	arg := UpdateEntrieParams{
		ID:     entry1.ID,
		Amount: sql.NullInt64{},
	}

	entry2, err := testQueries.UpdateEntrie(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.Equal(t, arg.Amount, entry2.Amount)
	require.WithinDuration(t, entry1.CreatedAt, entry2.CreatedAt, time.Second)
}

func TestQueries_DeleteEntrie(t *testing.T) {
	entry1 := CreateRandomEntry(t)
	err := testQueries.DeleteEntrie(context.Background(), entry1.ID)
	require.NoError(t, err)

	entry2, err := testQueries.GetEntrie(context.Background(), entry1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, entry2)
}

func TestQueries_ListEntries(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomEntry(t)
	}

	arg := ListEntriesParams{
		Limit:  5,
		Offset: 5,
	}

	entries, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, entries, 5)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
	}
}
