package types

const (
	// ModuleName defines the module name
	ModuleName = "treasury"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines message route key
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName
)

var (
	// global key of currency's denom list
	CurrenciesKey = []byte{0x11}
	// global key of currency state sequence
	SequenceKey = []byte{0x12}
)

var (
	// prefix for single currency key
	CurrencyPrefix = []byte{0x21}
)

// get the key for the currency from denom string
func GetCurrencyKey(denom string) []byte {
	return []byte(denom)
}
