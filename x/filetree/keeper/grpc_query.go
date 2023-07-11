package keeper

import (
	"github.com/MonikaCat/canine-chain/v2/x/filetree/types"
)

var _ types.QueryServer = Keeper{}
