package keeper

import (
	"github.com/MonikaCat/canine-chain/v2/x/storage/types"
)

var _ types.QueryServer = Keeper{}
