package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerr "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jackalLabs/canine-chain/x/storage/types"
)

func (k msgServer) PostContract(goCtx context.Context, msg *types.MsgPostContract) (*types.MsgPostContractResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, ok := k.GetProviders(ctx, msg.Creator)
	if !ok {
		return nil, fmt.Errorf("can't find provider")
	}

	contract, found := k.GetContracts(ctx, msg.Cid)
	if !found {
		return nil, sdkerr.Wrapf(sdkerr.ErrNotFound, "contract not found")
	}

	verified := VerifyDeal(contract.Merkle, msg.Hashlist, 0, msg.Item)
	if !verified {
		return nil, sdkerr.Wrapf(types.ErrCannotVerifyProof, "failed to verify proof")
	}

	validProvider := false

	providers := contract.Providers
	for _, provider := range providers {
		if provider == msg.Creator {
			validProvider = true
			break
		}
	}

	if !validProvider {
		return nil, sdkerr.Wrapf(sdkerr.ErrUnauthorized, "you are not listed as a provider on this deal!")
	}

	endBlock := ctx.BlockHeight() + contract.Duration
	if contract.Duration == 0 {
		endBlock = 0
	}

	deal := types.ActiveDeals{
		Cid:           contract.Cid,
		Signee:        contract.Creator,
		Provider:      msg.Creator,
		Startblock:    fmt.Sprintf("%d", ctx.BlockHeight()),
		Endblock:      fmt.Sprintf("%d", endBlock),
		Filesize:      fmt.Sprintf("%d", contract.Filesize),
		Proofverified: "true",
		Proofsmissed:  "0",
		Blocktoprove:  "0",
		Creator:       contract.Creator,
		Merkle:        contract.Merkle,
		Fid:           contract.Fid,
	}

	k.SetActiveDeals(ctx, deal)

	return &types.MsgPostContractResponse{}, nil
}
