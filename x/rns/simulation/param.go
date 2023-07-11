package simulation

import (
	"fmt"
	"math/rand"

	"github.com/MonikaCat/canine-chain/v2/x/rns/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

func ParamChanges(r *rand.Rand) []simtypes.ParamChange {
	_ = r
	return []simtypes.ParamChange{
		simulation.NewSimParamChange(types.ModuleName, string(types.KeyDepositAccount),
			func(r *rand.Rand) string {
				return fmt.Sprintf("\"%s\"", "jkl14c3j672kvw9l5uleh4x9uds2fre5vl7yun4ss8")
			},
		),
	}
}
