package cli

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"

	"github.com/rizon-world/rizon/x/treasury/types"
)

// NewQueryCmd returns the cli query commands for this module
func NewQueryCmd() *cobra.Command {
	// Group treasury queries under a subcommand
	treasuryQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the treasury module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	treasuryQueryCmd.AddCommand(
		GetCmdQueryCurrencies(),
		GetCmdQueryCurrency(),
		GetCmdQueryParams(),
	)

	return treasuryQueryCmd
}

// GetCmdQueryCurrencies queries all supported currency denom list
func GetCmdQueryCurrencies() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "currencies",
		Short: "Query all supported currency's denoms",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			pagination, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}
			params := &types.QueryCurrenciesRequest{
				Pagination: pagination,
			}

			res, err := queryClient.Currencies(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "all currencies")

	return cmd
}

// GetCmdQueryCurrency queries an information of single currency
func GetCmdQueryCurrency() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "currency [denom]",
		Short: "Query an information of single currency",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryCurrencyRequest{
				Denom: args[0],
			}

			res, err := queryClient.Currency(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res.Currency)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetCmdQueryParams queries the parameters of treasury module
func GetCmdQueryParams() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "params",
		Short: "Query the current parameters of treasury",
		Args:  cobra.NoArgs,
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
