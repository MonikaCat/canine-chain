package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/x/jklmint/types"
	storeTypes "github.com/jackalLabs/canine-chain/x/storage/types"
)

func (k Keeper) BlockMint(ctx sdk.Context) {
	mintTokens := sdk.NewDec(6_000_000)
	denom := k.GetParams(ctx).MintDenom

	providerRatio := sdk.NewDec(4)
	providerRatio = providerRatio.QuoInt64(10)
	validatorRatio := sdk.NewDec(6)
	validatorRatio = validatorRatio.QuoInt64(10)

	// get correct ratio
	providerTokens := mintTokens.Mul(providerRatio)
	validatorTokens := mintTokens.Mul(validatorRatio)

	// mint provider coins, update supply
	provCount := providerTokens.TruncateInt()
	providerCoin := sdk.NewCoin(denom, provCount)
	providerCoins := sdk.NewCoins(providerCoin)

	// mint validator coins, update supply
	valCount := validatorTokens.TruncateInt()
	validatorCoin := sdk.NewCoin(denom, valCount)
	validatorCoins := sdk.NewCoins(validatorCoin)

	totalCount := provCount.Add(valCount)
	// mint coins, update supply
	totalIntCoin := sdk.NewCoin(denom, totalCount)
	mintedCoin := totalIntCoin
	mintedCoins := sdk.NewCoins(mintedCoin)
	err := k.MintCoins(ctx, mintedCoins)
	if err != nil {
		panic(err)
	}

	err = k.bankKeeper.SendCoinsFromModuleToModule(ctx, types.ModuleName, storeTypes.ModuleName, providerCoins)
	if err != nil {
		panic(err)
	}

	// send the minted validator coins to the fee collector account
	err = k.AddCollectedFees(ctx, validatorCoins)
	if err != nil {
		panic(err)
	}

	// alerting network on mint amount
	defer telemetry.ModuleSetGauge(types.ModuleName, float32(totalCount.Int64()), "minted_tokens")

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeMint,
			sdk.NewAttribute(sdk.AttributeKeyAmount, fmt.Sprintf("%d", totalCount.Int64())),
		),
	)
}
