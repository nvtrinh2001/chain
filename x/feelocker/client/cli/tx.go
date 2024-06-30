package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"

	"github.com/bandprotocol/chain/v2/x/feelocker/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	flagPayees      = "payees"
	flagFee         = "fee"
	flagStatus      = "status"
	flagLockedFeeId = "locked-fee-id"
)

// NewTxCmd returns the transaction commands for this module
func NewTxCmd() *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "feelocker transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	txCmd.AddCommand(
		GetCmdLock(),
		GetCmdClaim(),
	)

	return txCmd
}

// GetCmdLock implements the lock command handler.
func GetCmdLock() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "lock --fee [coin] --payees [address1 address2 address3]",
		Short: "lock command",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			coinStr, err := cmd.Flags().GetString(flagFee)
			if err != nil {
				return err
			}

			fee, err := sdk.ParseCoinsNormalized(coinStr)
			if err != nil {
				return err
			}

			payeeStrList, err := cmd.Flags().GetStringSlice(flagPayees)
			if err != nil {
				return err
			}

			var payeeList []sdk.AccAddress
			for _, payee := range payeeStrList {
				payeeAccount, err := sdk.AccAddressFromBech32(payee)
				if err != nil {
					return err
				}
				payeeList = append(payeeList, payeeAccount)
			}

			msg := types.NewMsgLockRequest(clientCtx.GetFromAddress(), fee, payeeList)

			err = msg.ValidateBasic()
			if err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)

		},
	}

	// Add flags for payerAddress, payees, and fee
	cmd.Flags().StringSlice(flagPayees, []string{}, "List of payees")
	cmd.Flags().String(flagFee, "", "Fee")

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetCmdClaim implements the lock command handler.
func GetCmdClaim() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "claim --lodked-fee-id [id]",
		Short: "claim command",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			lockedFeeId, err := cmd.Flags().GetUint64(flagLockedFeeId)
			if err != nil {
				return err
			}

			msg := types.NewMsgClaimRequest(clientCtx.GetFromAddress(), types.LockedFeeID(lockedFeeId))

			err = msg.ValidateBasic()
			if err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)

		},
	}

	// Add flags for payerAddress, payees, and fee
	cmd.Flags().Uint64(flagLockedFeeId, 0, "Guarded Fee ID")

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
