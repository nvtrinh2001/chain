package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewDataSource(
	owner sdk.AccAddress, name, description, filename string, fee sdk.Coins, treasury sdk.AccAddress, id uint64,
) DataSource {
	return DataSource{
		Owner:             owner.String(),
		Name:              name,
		Description:       description,
		Filename:          filename,
		Treasury:          treasury.String(),
		Fee:               fee,
		RequirementFileId: id,
	}
}
