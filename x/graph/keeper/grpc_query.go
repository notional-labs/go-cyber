package keeper

import (
	"context"

	// "github.com/cosmos/cosmos-sdk/types/query"
	// "google.golang.org/grpc/codes"
	// "google.golang.org/grpc/status"
	"github.com/cybercongress/go-cyber/x/graph/types"

	// "github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ types.QueryServer = GraphKeeper{}

func (gk GraphKeeper) GraphStats(goCtx context.Context, _ *types.QueryGraphStatsRequest) (*types.QueryGraphStatsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	links := gk.GetLinksCount(ctx)
	cids := gk.GetCidsCount(ctx)
	return &types.QueryGraphStatsResponse{Cyberlinks: links, Particles: cids}, nil
}
