package bandwidth

import (
	"time"

	"github.com/cybercongress/go-cyber/x/bandwidth/keeper"
	"github.com/cybercongress/go-cyber/x/bandwidth/types"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var accountsToUpdate = make([]sdk.AccAddress, 0)

// recover and update bandwidth for accounts with changed stake
func updateAccountsMaxBandwidth(ctx sdk.Context, meter *keeper.BandwidthMeter) {
	for _, addr := range accountsToUpdate {
		meter.UpdateAccountMaxBandwidth(ctx, addr)
	}
	accountsToUpdate = make([]sdk.AccAddress, 0)
}

// collect all addresses with updated stake
func CollectAddressesWithStakeChange() func(ctx sdk.Context, addresses []sdk.AccAddress) {
	return func(ctx sdk.Context, addresses []sdk.AccAddress) {
		if ctx.IsCheckTx() {
			return
		}
		accountsToUpdate = append(accountsToUpdate, addresses...)
	}
}

func EndBlocker(ctx sdk.Context, bm *keeper.BandwidthMeter) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyEndBlocker)

	params := bm.GetParams(ctx)
	if ctx.BlockHeight() != 0 && uint64(ctx.BlockHeight())%params.AdjustPricePeriod == 0 {
		bm.AdjustPrice(ctx) // TODO Add block event for price and load
	}

	bm.CommitBlockBandwidth(ctx) // TODO add block event for committed bandwidth

	updateAccountsMaxBandwidth(ctx, bm)
}
