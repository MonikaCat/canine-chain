package wasmbinding

import (
	"encoding/json"
	"fmt"

	errorsmod "cosmossdk.io/errors"
	wasmvmtypes "github.com/CosmWasm/wasmvm/types"
	wasmkeeper "github.com/JackalLabs/jackal-wasmd/x/wasm/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"

	tx "github.com/cosmos/cosmos-sdk/types/tx"

	"github.com/jackalLabs/canine-chain/wasmbinding/bindings"
	filetreekeeper "github.com/jackalLabs/canine-chain/x/filetree/keeper"
	filetreetypes "github.com/jackalLabs/canine-chain/x/filetree/types"
	storagekeeper "github.com/jackalLabs/canine-chain/x/storage/keeper"

	testutils "github.com/jackalLabs/canine-chain/testutil"

	spew "github.com/davecgh/go-spew/spew"
)

// CustomMessageDecorator returns decorator for custom CosmWasm bindings messages
func CustomMessageDecorator(filetree *filetreekeeper.Keeper, storage *storagekeeper.Keeper) func(wasmkeeper.Messenger) wasmkeeper.Messenger {
	return func(old wasmkeeper.Messenger) wasmkeeper.Messenger {
		return &CustomMessenger{
			wrapped:  old,
			filetree: filetree,
			storage:  storage,
		}
	}
}

type CustomMessenger struct {
	wrapped  wasmkeeper.Messenger
	filetree *filetreekeeper.Keeper
	storage  *storagekeeper.Keeper
	handler  *SDKMessageHandler
}

var _ wasmkeeper.Messenger = (*CustomMessenger)(nil)

// DispatchMsg executes on the contractMsg.
func (m *CustomMessenger) DispatchMsg(ctx sdk.Context, contractAddr sdk.AccAddress, contractIBCPortID string, msg wasmvmtypes.CosmosMsg, sender string) ([]sdk.Event, [][]byte, error) {
	// sdkMsgs, err := m.handler.encoders.Encode(ctx, contractAddr, contractIBCPortID, msg)
	// if err != nil {
	// 	return nil, nil, err
	// }
	// logger, logFile := testutils.CreateLogger()

	// logger.Printf("The person who executed the contract is: %s", sender)

	// Please tell me it's possible to get the signer this way O.o?
	// for _, sdkMsg := range sdkMsgs {
	// 	logger.Println(sdkMsg.GetSigners())
	// }

	// logFile.Close()

	if msg.Custom != nil {
		// only handle the happy path where this is really posting files
		// leave everything else for the wrapped version
		var contractMsg bindings.JackalMsg
		if err := json.Unmarshal(msg.Custom, &contractMsg); err != nil {
			return nil, nil, errorsmod.Wrap(err, "Jackal msg")
		}
		if contractMsg.MakeRoot != nil {
			logger, logFile := testutils.CreateLogger()

			// We forked wasmd to get the bech32 address of the user who executed the contract :D
			// We do some verification here for security
			// Perhaps we can take this up a notch and make wasmd consume the contract executor as sdk.AccAddress
			// and build a signature verifcation function? To investigate

			var makeRoot bindings.MakeRoot
			makeRoot = *contractMsg.MakeRoot
			logger.Printf("The make root struct contains: %+v", makeRoot)
			logger.Printf("The make root struct contains: %+v", contractMsg.MakeRoot)
			fmt.Printf("The make root struct contains: %+v", contractMsg.MakeRoot)
			fmt.Printf("The make root struct contains: %+v", makeRoot)
			spew.Dump(makeRoot)

			// All our attempts to log the MakeRoot struct have failed thus far. Not sure why.

			// creator := contractMsg.MakeRoot.Creator

			logger.Printf("The person who called execute on the contract is: %s", sender)
			// logger.Printf("The person declared as msg.Creator for MakeRoot is: %s", creator)

			logFile.Close()

			return m.makeRoot(ctx, contractAddr, contractMsg.MakeRoot, sender) // need this
		}
		if contractMsg.PostFiles != nil {
			return m.postFiles(ctx, contractAddr, contractMsg.PostFiles) // need this
		}

	}
	return m.wrapped.DispatchMsg(ctx, contractAddr, contractIBCPortID, msg, sender)
}

// makeRoot posts a Files struct on chain that serves as the root directory for the user's filetree
// it is the merklePath of 's'
func (m *CustomMessenger) makeRoot(ctx sdk.Context, contractAddr sdk.AccAddress, makeRoot *bindings.MakeRoot, sender string) ([]sdk.Event, [][]byte, error) {
	err := PerformMakeRoot(m.filetree, ctx, contractAddr, makeRoot, sender)
	if err != nil {
		return nil, nil, errorsmod.Wrap(err, "perform make root")
	}
	return nil, nil, nil
}

func DecodeTx(txBytes []byte) (tx.Tx, error) {
	var raw tx.Tx
	err := raw.XXX_Unmarshal(txBytes)
	// TO DO
	// properly error handle
	if err != nil {
		return raw, err
	}
	return raw, nil
}

// PerformMakeRoot is used with makeRoot to post a root Files struct to chain, as described above; validates the msgMakeRoot.
func PerformMakeRoot(f *filetreekeeper.Keeper, ctx sdk.Context, contractAddr sdk.AccAddress, makeRoot *bindings.MakeRoot, sender string) error {
	if makeRoot == nil {
		return wasmvmtypes.InvalidRequest{Err: "make root null make root"}
	}

	// This isn't need at the moment--especially because we can't seem to log the Creator address
	// if sender != makeRoot.Creator {
	// 	return wasmvmtypes.InvalidRequest{Err: "You can only create a root filetree File for yourself!"} // Desperately need better error handling
	// }

	// Silly me, we can just pass in the executor of the contract to make a root

	sdkMsg := filetreetypes.NewMsgMakeRootV2(
		sender,
		makeRoot.Editors,
		makeRoot.Viewers,
		makeRoot.TrackingNumber,
	)
	if err := sdkMsg.ValidateBasic(); err != nil {
		return err
	}

	msgServer := filetreekeeper.NewMsgServerImpl(*f)
	_, err := msgServer.MakeRootV2(sdk.WrapSDKContext(ctx), sdkMsg)
	if err != nil {
		return errorsmod.Wrap(err, "making root from message")
	}

	return nil
}

// postFiles posts a Files struct to chain
func (m *CustomMessenger) postFiles(ctx sdk.Context, contractAddr sdk.AccAddress, postFiles *bindings.PostFiles) ([]sdk.Event, [][]byte, error) {
	err := PerformPostFiles(m.filetree, ctx, contractAddr, postFiles)
	if err != nil {
		return nil, nil, errorsmod.Wrap(err, "perform post files")
	}
	return nil, nil, nil
}

// PerformPostFiles is used with postFiles to post a Files struct to chian; validates the msgPostFiles.
func PerformPostFiles(f *filetreekeeper.Keeper, ctx sdk.Context, contractAddr sdk.AccAddress, postFiles *bindings.PostFiles) error {
	if postFiles == nil {
		return wasmvmtypes.InvalidRequest{Err: "post files null post files"}
	}

	msgServer := filetreekeeper.NewMsgServerImpl(*f)

	msgPostFiles := filetreetypes.NewMsgPostFile(
		postFiles.Creator,
		postFiles.Account,
		postFiles.HashParent,
		postFiles.HashChild,
		postFiles.Contents,
		postFiles.Viewers,
		postFiles.Editors,
		postFiles.TrackingNumber,
	)

	if err := msgPostFiles.ValidateBasic(); err != nil {
		return errorsmod.Wrap(err, "failed validating MsgPostFiles")
	}

	// Post files
	_, err := msgServer.PostFile(
		sdk.WrapSDKContext(ctx),
		msgPostFiles,
	)
	if err != nil {
		return errorsmod.Wrap(err, "creating denom")
	}
	return nil
}

// Leave here because we might need this

// parseAddress parses address from bech32 string and verifies its format.
// func parseAddress(addr string) (sdk.AccAddress, error) {
// 	parsed, err := sdk.AccAddressFromBech32(addr)
// 	if err != nil {
// 		return nil, errorsmod.Wrap(err, "address from bech32")
// 	}
// 	err = sdk.VerifyAddressFormat(parsed)
// 	if err != nil {
// 		return nil, errorsmod.Wrap(err, "verify address format")
// 	}
// 	return parsed, nil
// }
