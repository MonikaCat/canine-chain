package keeper

import (
	"github.com/MonikaCat/canine-chain/v2/x/notifications/types"
)

var _ types.QueryServer = Keeper{}
