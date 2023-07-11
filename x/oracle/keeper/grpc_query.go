package keeper

import (
	"github.com/MonikaCat/canine-chain/v2/x/oracle/types"
)

var _ types.QueryServer = Keeper{}
