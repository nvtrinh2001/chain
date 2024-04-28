package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewRequirementFile(
	owner sdk.AccAddress, name, description, filename string, fee sdk.Coins, treasury sdk.AccAddress,
) RequirementFile {
	return RequirementFile{
		Owner:       owner.String(),
		Name:        name,
		Description: description,
		Filename:    filename,
		Treasury:    treasury.String(),
		Fee:         fee,
	}
}
