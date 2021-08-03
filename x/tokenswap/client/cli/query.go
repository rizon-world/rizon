package cli

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"

	"github.com/rizon-world/rizon/x/tokenswap/types"
)

// NewQueryCmd returns the cli query commands for this module
func NewQueryCmd() *cobra.Command {
	// Group tokenswap queries under a subcommand
	tokenswapQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the tokenswap module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	tokenswapQueryCmd.AddCommand(
		GetCmdQuerySwap(),
		GetCmdQuerySwappedAmount(),
		GetCmdQueryParams(),
	)

	return tokenswapQueryCmd
}

// GetCmdQuerySwap queries a tokenswap request by tx hash
func GetCmdQuerySwap() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get [tx-hash]",
		Short: "Query swap request by tx hash",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryTokenswapRequest{
				TxHash: args[0],
			}

			res, err := queryClient.Tokenswap(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res.Tokenswap)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetCmdQuerySwappedAmount queries current swapped amount of tokenswap module
func GetCmdQuerySwappedAmount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "swapped-amount",
		Args:  cobra.NoArgs,
		Short: "Query current swapped amount of tokenswap",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.SwappedAmount(context.Background(), &types.QuerySwappedAmountRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(&res.SwappedAmount)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetCmdQueryParams queries the parameters of tokenswap module
func GetCmdQueryParams() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "params",
		Args:  cobra.NoArgs,
		Short: "Query the current parameters of tokenswap",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.Params(context.Background(), &types.QueryParamsRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(&res.Params)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
