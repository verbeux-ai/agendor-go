package agendor_test

import (
	"context"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/verbeux-ai/agendor-go"
)

func TestCreatePeople(t *testing.T) {
	ctx := context.Background()
	people, err := client.CreatePeople(ctx, agendor.CreatePeopleRequest{
		Name: "Teste agendor-go",
	})
	require.NoError(t, err)
	require.NotEmpty(t, people)
}

func TestCreatePeopleDeal(t *testing.T) {
	ctx := context.Background()

	peopleIDStr := os.Getenv("PEOPLE_ID")
	require.NotEmpty(t, peopleIDStr)
	peopleID, err := strconv.Atoi(peopleIDStr)
	require.NoError(t, err)

	deal, err := client.CreatePeopleDeal(ctx, peopleID, agendor.CreatePeopleDealRequest{
		Title: "teste3",
	})
	require.NoError(t, err)
	require.NotEmpty(t, deal)
}

func TestListPeople(t *testing.T) {
	ctx := context.Background()
	peoples, err := client.ListPeoples(ctx, agendor.FilterPeoplesRequest{
		Name: "Teste",
	})
	require.NoError(t, err)
	require.NotEmpty(t, peoples)
}
