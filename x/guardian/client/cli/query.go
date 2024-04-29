package cli

import (
	"context"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	"github.com/bandprotocol/chain/v2/x/guardian/types"
)

// GetQueryCmd returns the cli query commands for this module.
func GetQueryCmd() *cobra.Command {
	guardianCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the guardian module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	guardianCmd.AddCommand(
		GetQueryCmdGuardedFee(),
		GetQueryCmdGuardedFeeList(),
	)
	return guardianCmd
}

func GetQueryCmdGuardedFee() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "guarded-fee [id]",
		Short: "Get summary information of a guarded fee",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)
			r, err := queryClient.QueryGuardedFee(context.Background(), &types.QueryGuardedFeeRequest{GuardedFeeId: id})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(r)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func GetQueryCmdGuardedFeeList() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "guarded-fee-list [account-address] --status [locked|claimable|claimed]",
		Short: "Get summary information of a guarded fee list based on an account address",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			statusStr, err := cmd.Flags().GetString(flagStatus)
			if err != nil {
				return err
			}

			var status types.STATUS
			switch statusStr {
			case "claimable":
				status = types.STATUS_CLAIMABLE
			case "claimed":
				status = types.STATUS_CLAIMED
			default:
				status = -1
			}

			queryClient := types.NewQueryClient(clientCtx)
			r, err := queryClient.QueryGuardedFeeList(context.Background(), &types.QueryGuardedFeeListRequest{AccountAddress: args[0], Status: status})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(r)
		},
	}
	cmd.Flags().String(flagStatus, "", "Status")
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
