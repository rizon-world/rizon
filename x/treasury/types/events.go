package types

// treasury event types and attribute keys
const (
	EventTypeMint           = "mint"
	EventTypeBurn           = "burn"
	EventTypeCurrencyUpdate = "currencyUpdate"

	AttributeKeyDenom      = "denom"
	AttributeKeyDesc       = "desc"
	AttributeKeyMintable   = "mintable"
	AttributeKeyOwner      = "owner"
	AttributeKeyReceiver   = "receiver"
	AttributeValueCategory = ModuleName
)
