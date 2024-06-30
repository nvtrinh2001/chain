package yoda

import (
	"fmt"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/bandprotocol/chain/v2/x/oracle/types"
)

type rawRequest struct {
	dataSourceID           types.DataSourceID
	dataSourceHash         string
	externalID             types.ExternalID
	calldata               string
	requirementFileID      types.RequirementFileID
	requirementFileHash    string
	offlineFeeLimit        sdk.Coins
	baseOffchainFeePerHour uint64
	language               string
	usedExternalLibraries  string
}

// GetRawRequests returns the list of all raw data requests in the given log.
func GetRawRequests(log sdk.ABCIMessageLog) ([]rawRequest, error) {
	dataSourceIDs := GetEventValues(log, types.EventTypeRawRequest, types.AttributeKeyDataSourceID)
	dataSourceHashList := GetEventValues(log, types.EventTypeRawRequest, types.AttributeKeyDataSourceHash)
	externalIDs := GetEventValues(log, types.EventTypeRawRequest, types.AttributeKeyExternalID)
	calldataList := GetEventValues(log, types.EventTypeRawRequest, types.AttributeKeyCalldata)
	requirementFileIDs := GetEventValues(log, types.EventTypeRawRequest, types.AttributeKeyRequirementFileID)
	requirementFileHashList := GetEventValues(log, types.EventTypeRawRequest, types.AttributeKeyRequirementFileHash)
	offlineFeeLimitList := GetEventValues(log, types.EventTypeRawRequest, types.AttributeOffchainFeeLimit)
	baseOffchainFeePerHourList := GetEventValues(log, types.EventTypeRawRequest, types.AttributeBaseOffchainFeePerHour)
	languageList := GetEventValues(log, types.EventTypeRawRequest, types.AttributeKeyLanguage)
	usedExternalLibraryList := GetEventValues(log, types.EventTypeRawRequest, types.AttributeKeyUsedExternalLibraries)

	if len(dataSourceIDs) != len(externalIDs) {
		return nil, fmt.Errorf("Inconsistent data source count and external ID count")
	}
	if len(dataSourceIDs) != len(calldataList) {
		return nil, fmt.Errorf("Inconsistent data source count and calldata count")
	}

	var reqs []rawRequest
	for idx := range dataSourceIDs {
		dataSourceID, err := strconv.Atoi(dataSourceIDs[idx])
		if err != nil {
			return nil, fmt.Errorf("Failed to parse data source id: %s", err.Error())
		}

		externalID, err := strconv.Atoi(externalIDs[idx])
		if err != nil {
			return nil, fmt.Errorf("Failed to parse external id: %s", err.Error())
		}

		requirementFileID, err := strconv.Atoi(requirementFileIDs[idx])
		if err != nil {
			return nil, fmt.Errorf("Failed to parse requirement file id: %s", err.Error())
		}

		offlineFeeLimit, err := sdk.ParseCoinsNormalized(offlineFeeLimitList[idx])
		if err != nil {
			return nil, err
		}

		baseOffchainFeePerHour, err := strconv.Atoi(baseOffchainFeePerHourList[idx])
		if err != nil {
			return nil, err
		}

		language := languageList[idx]
		usedExternalLibraries := usedExternalLibraryList[idx]

		reqs = append(reqs, rawRequest{
			dataSourceID:           types.DataSourceID(dataSourceID),
			dataSourceHash:         dataSourceHashList[idx],
			externalID:             types.ExternalID(externalID),
			calldata:               calldataList[idx],
			requirementFileID:      types.RequirementFileID(requirementFileID),
			requirementFileHash:    requirementFileHashList[idx],
			offlineFeeLimit:        offlineFeeLimit,
			baseOffchainFeePerHour: uint64(baseOffchainFeePerHour),
			language:               language,
			usedExternalLibraries:  usedExternalLibraries,
		})
	}
	return reqs, nil
}

// GetEventValues returns the list of all values in the given log with the given type and key.
func GetEventValues(log sdk.ABCIMessageLog, evType string, evKey string) (res []string) {
	for _, ev := range log.Events {
		if ev.Type != evType {
			continue
		}

		for _, attr := range ev.Attributes {
			if attr.Key == evKey {
				res = append(res, attr.Value)
			}
		}
	}
	return res
}

// GetEventValue checks and returns the exact value in the given log with the given type and key.
func GetEventValue(log sdk.ABCIMessageLog, evType string, evKey string) (string, error) {
	values := GetEventValues(log, evType, evKey)
	if len(values) == 0 {
		return "", fmt.Errorf("Cannot find event with type: %s, key: %s", evType, evKey)
	}
	if len(values) > 1 {
		return "", fmt.Errorf("Found more than one event with type: %s, key: %s", evType, evKey)
	}
	return values[0], nil
}
