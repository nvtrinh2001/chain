package types

// DefaultGenesisState returns the default oracle genesis state.
func DefaultGenesisState() *GenesisState {
	return &GenesisState{}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (g GenesisState) Validate() error { return nil }
