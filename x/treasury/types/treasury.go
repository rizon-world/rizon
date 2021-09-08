package types

// NewCurrency creates new currency
func NewCurrency(denom string, desc string, owner string, mintable bool) Currency {
	return Currency{
		Denom:    denom,
		Desc:     desc,
		Owner:    owner,
		Mintable: mintable,
	}
}

// NewCurrencies creates new currency list
func NewCurrencies(denoms []string) Currencies {
	return Currencies{
		Denoms: denoms,
	}
}

// NewMaxAtoloSupply creates a new max atolo supply
func NewMaxAtoloSupply(amount int64) MaxAtoloSupply {
	return MaxAtoloSupply{
		Amount: amount,
	}
}

// NewSequence creates new sequence
func NewSequence(number int64) Sequence {
	return Sequence{
		Number: number,
	}
}
