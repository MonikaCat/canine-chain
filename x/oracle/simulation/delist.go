package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/jackalLabs/canine-chain/v3/x/rns/keeper"
	"github.com/jackalLabs/canine-chain/v3/x/rns/types"
)

func SimulateMsgDelist(
	_ types.AccountKeeper,
	_ types.BankKeeper,
	_ keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgDelist{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the Delist simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "Delist simulation not implemented"), nil, nil
	}
}
