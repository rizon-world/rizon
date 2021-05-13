package types

import (
	"fmt"

	yaml "gopkg.in/yaml.v2"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

const (
	// treasury module is enabled or not
	DefaultMintable = true
	DefaultOwner    = "rizon136fzkc73rm5def5fngs386qdlxcuvxvrte8lk7"
	DefaultSequence = 1
)

var (
	DefaultCurrencyList = []Currency{
		{
			Denom:    "skrw",
			Desc:     "stable coin of KRW",
			Owner:    DefaultOwner,
			Mintable: true,
		},
		{ // susd is for test
			Denom:    "susd",
			Desc:     "stable coin of USD",
			Owner:    DefaultOwner,
			Mintable: true,
		},
	}
)
var (
	KeyMintable      = []byte("Mintable")
	KeyCurrencyList  = []byte("CurrencyList")
	KeyParamSequence = []byte("ParamSequence")
)

var (
	_ paramtypes.ParamSet = (*Params)(nil)
)

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return Params{
		Mintable:     DefaultMintable,
		CurrencyList: DefaultCurrencyList,
		Sequence:     DefaultSequence,
	}
}

// implements params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyMintable, &p.Mintable, validateMintable),
		paramtypes.NewParamSetPair(KeyCurrencyList, &p.CurrencyList, validateCurrencyList),
		paramtypes.NewParamSetPair(KeyParamSequence, &p.Sequence, validateSequence),
	}
}

// String returns a human readable string representation of the parameters
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

// making keyTable for treasury module
func ParamsKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// validate a set of params
func (p *Params) Validate() error {
	if err := validateMintable(p.Mintable); err != nil {
		return err
	}

	if err := validateCurrencyList(p.CurrencyList); err != nil {
		return err
	}

	if err := validateSequence(p.Sequence); err != nil {
		return err
	}

	return nil
}

// validate mintable
func validateMintable(i interface{}) error {
	_, ok := i.(bool)
	if !ok {
		return fmt.Errorf("invalid mintable parameter type: %T", i)
	}

	return nil
}

// validate currency list
func validateCurrencyList(i interface{}) error {
	v, ok := i.([]Currency)
	if !ok {
		return fmt.Errorf("invalid currency list parameter type: %T", i)
	}

	for _, d := range v {
		owner, err := sdk.AccAddressFromBech32(d.Owner)
		if err != nil {
			return err
		}

		if owner == nil || owner.Empty() {
			return fmt.Errorf("owner should not nil or empty")
		}
	}

	return nil
}

// validate sequence
func validateSequence(i interface{}) error {
	_, ok := i.(int64)
	if !ok {
		return fmt.Errorf("invalid sequence parameter type: %T", i)
	}

	return nil
}
