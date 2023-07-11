package cli

import (
	"context"

	"github.com/MonikaCat/canine-chain/v2/x/storage/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func CmdGetStorageStats() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-storage-stats",
		Short: "lists stats about storage on the network",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryStorageStatsRequest{}

			res, err := queryClient.StorageStats(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
