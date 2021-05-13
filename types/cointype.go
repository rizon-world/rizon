package types

const (
	// CoinType is the ATOLO coin type as defined in SLIP44 (https://github.com/satoshilabs/slips/blob/master/slip-0044.md)
	CoinType = 1217

	// FullFundraiserPath is the parts of the BIP44 HD path that are fixed by
	// what we used during the ATOLO fundraiser.
	FullFundraiserPath = "m/44'/1217'/0'/0/0"

	// DefaultDenom is base cointype of ATOLO coin
	DefaultDenom = "uatolo"
)
