package cli

import (
	"strconv"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/rizon-world/rizon/x/tokenswap/types"
)

// NewTxCmd returns the transaction commands for this module
func NewTxCmd() *cobra.Command {
	tokenswapTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "tokenswap transactions subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	tokenswapTxCmd.AddCommand(NewCreateSwapCmd())

	return tokenswapTxCmd
}

// NewCreateSwapCmd is the CLI command for creating a token swap request
func NewCreateSwapCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create [tx-hash] [amount] [receiver-address] --from [address]",
		Short: "Create a token swap request",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			amountArg, err := strconv.ParseInt(args[1], 10, 64)
			if err != nil {
				return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "amount must be an integer")
			}
			amount := sdk.NewDec(amountArg)

			_, err = sdk.AccAddressFromBech32(args[2])
			if err != nil {
				return err
			}

			signer := clientCtx.GetFromAddress()

			msg := types.NewMsgCreateTokenswapRequest(
				args[0],
				args[2],
				signer.String(),
				amount,
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
