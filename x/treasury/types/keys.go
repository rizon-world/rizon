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
	KeyCurrencies = []byte{0x11}
	// global key of currency state sequence
	KeySequence = []byte{0x12}
)

var (
	// prefix for single currency key
	PrefixCurrency = []byte{0x21}
)
