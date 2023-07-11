package simulation

import (
	"math/rand"

	"github.com/MonikaCat/canine-chain/v2/x/rns/keeper"
	"github.com/MonikaCat/canine-chain/v2/x/rns/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgCancelBid(
	_ types.AccountKeeper,
	_ types.BankKeeper,
	_ keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgCancelBid{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the CancelBid simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "CancelBid simulation not implemented"), nil, nil
	}
}
