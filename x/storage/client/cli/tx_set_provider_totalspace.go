package cli

import (
	"strconv"

	"github.com/MonikaCat/canine-chain/v2/x/storage/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdSetProviderTotalspace() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-totalspace [space]",
		Short: "Broadcast message set-provider-totalspace",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argSpace := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSetProviderTotalspace(
				clientCtx.GetFromAddress().String(),
				argSpace,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	return cmd
}
