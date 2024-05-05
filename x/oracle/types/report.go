package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewReport(
	Validator sdk.ValAddress,
	InBeforeResolve bool,
	RawReports []RawReport,
	OffchainFeeUsed sdk.Coins,
) Report {
	return Report{
		Validator:       Validator.String(),
		InBeforeResolve: InBeforeResolve,
		RawReports:      RawReports,
		OffchainFeeUsed: OffchainFeeUsed,
	}
}

func NewRawReport(
	ExternalID ExternalID,
	ExitCode uint32,
	Data []byte,
	OffchainFeeUsed sdk.Coins,
) RawReport {
	return RawReport{
		ExternalID:      ExternalID,
		ExitCode:        ExitCode,
		Data:            Data,
		OffchainFeeUsed: OffchainFeeUsed,
	}
}
