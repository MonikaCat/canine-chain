package simulation

import (
	"github.com/MonikaCat/canine-chain/v2/x/rns/types"
	"github.com/cosmos/cosmos-sdk/types/module"
)

func RandomizedGenState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}

	rnsGenesis := types.DefaultGenesis()
	p := types.DefaultParams()
	p.DepositAccount = "jkl1arsaayyj5tash86mwqudmcs2fd5jt5zgc3sexc"
	rnsGenesis.Params = p

	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(rnsGenesis)
}
