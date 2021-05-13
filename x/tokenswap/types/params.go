package types

import (
	"fmt"

	yaml "gopkg.in/yaml.v2"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

const (
	// tokenswap module is enabled or not
	DefaultSwappable = true
	// who can request tokenswap
	DefaultSigner = "rizon136fzkc73rm5def5fngs386qdlxcuvxvrte8lk7"
)

var (
	KeySwappable = []byte("Swappable")
	KeySigner    = []byte("Signer")
)

var (
	_ paramtypes.ParamSet = (*Params)(nil)
)

// NewParams creates a new Params instance
func NewParams(swappable bool, signer string) Params {
	return Params{
		Swappable: swappable,
		Signer:    signer,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return Params{
		Swappable: DefaultSwappable,
		Signer:    DefaultSigner,
	}
}

// implements params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeySwappable, &p.Swappable, validateSwappable),
		paramtypes.NewParamSetPair(KeySigner, &p.Signer, validateSigner),
	}
}

// String returns a human readable string representation of the parameters
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

// making keyTable for tokenswap module
func ParamsKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// validate a set of params
func (p *Params) Validate() error {
	if err := validateSwappable(p.Swappable); err != nil {
		return err
	}

	if err := validateSigner(p.Signer); err != nil {
		return err
	}

	return nil
}

// validate swappable
func validateSwappable(i interface{}) error {
	_, ok := i.(bool)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return nil
}

// validate signer
func validateSigner(i interface{}) error {
	v, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	signer, err := sdk.AccAddressFromBech32(v)
	if err != nil {
		return err
	}

	if signer == nil || signer.Empty() {
		return fmt.Errorf("signer should not nil or empty")
	}
	return nil
}
