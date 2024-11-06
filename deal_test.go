package agendor_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/verbeux-ai/agendor-go"
)

func TestUpdateDeal(t *testing.T) {
	ctx := context.Background()
	deal, err := client.UpdateDeal(ctx, 27484945, agendor.UpdateDealRequest{
		Description: "Teste",
	})
	require.NoError(t, err)
	require.NotEmpty(t, deal)
}

func TestUpdateDealStatus(t *testing.T) {
	ctx := context.Background()
	deal, err := client.UpdateDealStatus(ctx, 27484945, agendor.UpdateDealStatusRequest{
		DealStatusText: agendor.DealStatusOngoing,
	})
	require.NoError(t, err)
	require.NotEmpty(t, deal)
}

func TestClient_UpdateDealStage(t *testing.T) {
	ctx := context.Background()
	deal, err := client.UpdateDealStage(ctx, 27484945, agendor.UpdateDealStageRequest{
		DealStage: 1,
	})
	require.NoError(t, err)
	require.NotEmpty(t, deal)
}
