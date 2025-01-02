package sqlc_test

import (
	"context"
	"github.com/stretchr/testify/require"
	"readly/db/sqlc"
	"readly/testdata"
	"testing"
)

func createRandomPublisher(t *testing.T) db.Publisher {
	arg := testdata.RandomString(6)
	publisher, err := querier.CreatePublisher(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, publisher)
	return publisher
}

func TestCreatePublisher(t *testing.T) {
	createRandomPublisher(t)
}

func TestGetPublisherByName(t *testing.T) {
	publisher1 := createRandomPublisher(t)
	publisher2, err := querier.GetPublisherByName(context.Background(), publisher1.Name)
	require.NoError(t, err)
	require.NotEmpty(t, publisher2)
	require.Equal(t, publisher1.Name, publisher2.Name)
}

func TestDeletePublisher(t *testing.T) {
	publisher1 := createRandomPublisher(t)
	err := querier.DeletePublisher(context.Background(), publisher1.Name)
	require.NoError(t, err)

	publisher2, err := querier.GetPublisherByName(context.Background(), publisher1.Name)
	require.Error(t, err)
	require.Empty(t, publisher2)
}

func TestGetAllPublishers(t *testing.T) {
	for i := 0; i < 4; i++ {
		createRandomPublisher(t)
	}

	arg := db.GetAllPublishersParams{
		Limit:  2,
		Offset: 0,
	}

	publishers, err := querier.GetAllPublishers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, publishers, 2)

	for _, publisher := range publishers {
		require.NotEmpty(t, publisher)
	}
}
