package cli

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/MonikaCat/canine-chain/v2/x/notifications/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdBlockSenders() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "block-senders [sender-ids]",
		Short: "Broadcast message BlockSenders",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argSenderIds := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			senderAddresses := strings.Split(argSenderIds, ",")

			var senderIds []string

			for _, v := range senderAddresses {
				if len(v) < 1 {
					continue
				}
				senderIds = append(senderIds, v)
			}

			jsonSenders, err := json.Marshal(senderIds)
			if err != nil {
				return err
			}

			msg := types.NewMsgBlockSenders(
				clientCtx.GetFromAddress().String(),
				string(jsonSenders),
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
