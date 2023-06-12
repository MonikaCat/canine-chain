package keeper

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"strconv"

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

	cids := []string{contract.Cid}

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

	for i := 0; i < 2; i++ {
		h := sha256.New()
		_, err := io.WriteString(h, fmt.Sprintf("%s%d", contract.Cid, i))
		if err != nil {
			return nil, err
		}
		hashName := h.Sum(nil)

		scid, err := MakeCid(hashName)
		if err != nil {
			return nil, err
		}

		newContract := types.Strays{
			Cid:      scid,
			Signee:   contract.Creator,
			Fid:      contract.Fid,
			Filesize: strconv.FormatUint(contract.Filesize, 10),
			Merkle:   contract.Merkle,
			End:      endBlock,
		}

		cids = append(cids, scid)

		k.SetStrays(ctx, newContract)

	}

	cidArr, err := json.Marshal(cids)
	if err != nil {
		return nil, err
	}

	newFtoC := types.FidCid{
		Fid:  contract.Fid,
		Cids: string(cidArr),
	}
	k.SetFidCid(ctx, newFtoC)

	return &types.MsgPostContractResponse{}, nil
}
