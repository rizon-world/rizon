package types

const (
	// ModuleName defines the module name
	ModuleName = "tokenswap"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines message route key
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName
)

var (
	// global key of current swapped amount
	KeySwappedAmount = []byte{0x11}
)
