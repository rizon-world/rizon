package types

// query endpoints supported by the Querier
const (
	QueryTokenswap = "get"
	QueryParams    = "parameters"
)

// QueryTokenswapParam defines the param for the following query:
// - 'custom/tokenswap/get/'
type QueryTokenswapParam struct {
	TxHash string
}

// NewQueryTokenswapParam creates a new QueryTokenswapParam
func NewQueryTokenswapParam(txHash string) QueryTokenswapParam {
	return QueryTokenswapParam{
		TxHash: txHash,
	}
}
