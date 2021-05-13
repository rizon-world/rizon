package types

// query endpoints supported by the Querier
const (
	QueryCurrencies = "currencies"
	QueryCurrency   = "currency"
	QueryParams     = "parameters"
)

// QueryCurrencyParam defines the param for the following query:
// - 'custom/treasury/currencies/{denom}'
type QueryCurrencyParam struct {
	Denom string
}

// NewQueryCurrencyParam creates a new QueryCurrencyParam
func NewQueryCurrencyParam(denom string) QueryCurrencyParam {
	return QueryCurrencyParam{
		Denom: denom,
	}
}
