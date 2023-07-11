package keeper

import (
	"github.com/MonikaCat/canine-chain/v2/x/rns/types"
)

var _ types.QueryServer = Keeper{}
