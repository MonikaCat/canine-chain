package keeper

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/MonikaCat/canine-chain/v2/x/filetree/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) AddViewers(goCtx context.Context, msg *types.MsgAddViewers) (*types.MsgAddViewersResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	file, found := k.GetFiles(ctx, msg.Address, msg.Fileowner)
	if !found {
		return nil, types.ErrFileNotFound
	}

	isOwner := IsOwner(file, msg.Creator)
	if !isOwner {
		return nil, types.ErrCannotAllowView
	}

	pvacc := file.ViewingAccess

	jvacc := make(map[string]string)
	if err := json.Unmarshal([]byte(pvacc), &jvacc); err != nil {
		return nil, types.ErrCantUnmarshall
	}

	ids := strings.Split(msg.ViewerIds, ",")
	keys := strings.Split(msg.ViewerKeys, ",")

	for i, v := range ids {
		jvacc[v] = keys[i]
	}

	vaccbytes, err := json.Marshal(jvacc)
	if err != nil {
		return nil, types.ErrCantMarshall
	}
	newviewers := string(vaccbytes)

	file.ViewingAccess = newviewers

	k.SetFiles(ctx, file)

	return &types.MsgAddViewersResponse{}, nil
}
