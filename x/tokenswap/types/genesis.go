package types

// NewGenesisState creates a new GenesisState instance
func NewGenesisState(params Params) *GenesisState {
	return &GenesisState{
		Params: params,
	}
}

// Validate checks all genesis states are valid
func (s *GenesisState) Validate() error {
	if err := s.Params.Validate(); err != nil {
		return err
	}

	return nil
}

// DefaultGenesisState returns default genesis state of tokenswap module
func DefaultGenesisState() *GenesisState {
	return NewGenesisState(DefaultParams())
}
