package yoda

import (
	"encoding/hex"
	"github.com/bandprotocol/chain/v2/yoda/executor"
	"strconv"

	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/crypto/tmhash"

	"github.com/bandprotocol/chain/v2/x/oracle/types"
)

type processingResult struct {
	rawReport types.RawReport
	version   string
	err       error
}

func MustAtoi(num string) int64 {
	result, err := strconv.ParseInt(num, 10, 64)
	if err != nil {
		panic(err)
	}
	return result
}

func handleTransaction(c *Context, l *Logger, tx abci.TxResult) {
	l.Debug(":eyes: Inspecting incoming transaction: %X", tmhash.Sum(tx.Tx))
	if tx.Result.Code != 0 {
		l.Debug(":alien: Skipping transaction with non-zero code: %d", tx.Result.Code)
		return
	}

	logs, err := sdk.ParseABCILogs(tx.Result.Log)
	if err != nil {
		l.Error(":cold_sweat: Failed to parse transaction logs with error: %s", c, err.Error())
		return
	}

	for _, log := range logs {
		go handleRequestLog(c, l, log)
	}
}

func handleRequestLog(c *Context, l *Logger, log sdk.ABCIMessageLog) {
	idStr, err := GetEventValue(log, types.EventTypeRequest, types.AttributeKeyID)
	if err != nil {
		l.Debug(":cold_sweat: Failed to parse request id with error: %s", err.Error())
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		l.Error(":cold_sweat: Failed to convert %s to integer with error: %s", c, idStr, err.Error())
		return
	}

	l = l.With("rid", id)

	// If id is in pending requests list, then skip it.
	if c.pendingRequests[types.RequestID(id)] {
		l.Debug(":eyes: Request is in pending list, then skip")
		return
	}

	// Skip if not related to this validator
	validators := GetEventValues(log, types.EventTypeRequest, types.AttributeKeyValidator)
	hasMe := false
	for _, validator := range validators {
		if validator == c.validator.String() {
			hasMe = true
			break
		}
	}

	if !hasMe {
		l.Debug(":next_track_button: Skip request not related to this validator")
		return
	}

	l.Info(":delivery_truck: Processing incoming request event")

	reqs, err := GetRawRequests(log)
	if err != nil {
		l.Error(":skull: Failed to parse raw requests with error: %s", c, err.Error())
	}

	keyIndex := c.nextKeyIndex()
	key := c.keys[keyIndex]

	reports, execVersions := handleRawRequests(c, l, types.RequestID(id), reqs, key)

	rawAskCount := GetEventValues(log, types.EventTypeRequest, types.AttributeKeyAskCount)
	if len(rawAskCount) != 1 {
		panic("Fail to get ask count")
	}
	askCount := MustAtoi(rawAskCount[0])

	rawMinCount := GetEventValues(log, types.EventTypeRequest, types.AttributeKeyMinCount)
	if len(rawMinCount) != 1 {
		panic("Fail to get min count")
	}
	minCount := MustAtoi(rawMinCount[0])

	rawCallData := GetEventValues(log, types.EventTypeRequest, types.AttributeKeyCalldata)
	if len(rawCallData) != 1 {
		panic("Fail to get call data")
	}
	callData, err := hex.DecodeString(rawCallData[0])
	if err != nil {
		l.Error(":skull: Fail to parse call data: %s", c, err.Error())
	}

	var clientID string
	rawClientID := GetEventValues(log, types.EventTypeRequest, types.AttributeKeyClientID)
	if len(rawClientID) > 0 {
		clientID = rawClientID[0]
	}

	offchainFeeUsed := reports[0].OffchainFeeUsed
	for idx := 1; idx < len(reports); idx++ {
		for _, coin := range reports[idx].OffchainFeeUsed {
			for i, offchainCoin := range offchainFeeUsed {
				if offchainCoin.Denom == coin.Denom {
					offchainFeeUsed[i].Amount.Add(coin.Amount)
					break
				}
			}
		}
	}

	c.pendingMsgs <- ReportMsgWithKey{
		msg:         types.NewMsgReportData(types.RequestID(id), reports, c.validator, offchainFeeUsed),
		execVersion: execVersions,
		keyIndex:    keyIndex,
		feeEstimationData: FeeEstimationData{
			askCount:    askCount,
			minCount:    minCount,
			callData:    callData,
			rawRequests: reqs,
			clientID:    clientID,
		},
	}
}

func handlePendingRequest(c *Context, l *Logger, id types.RequestID) {
	req, err := GetRequest(c, l, id)
	if err != nil {
		l.Error(":skull: Failed to get request with error: %s", c, err.Error())
		return
	}

	l.Info(":delivery_truck: Processing pending request")

	keyIndex := c.nextKeyIndex()
	key := c.keys[keyIndex]

	var rawRequests []rawRequest

	// prepare raw requests
	for _, raw := range req.RawRequests {
		hash, err := GetDataSourceHash(c, l, raw.DataSourceID)
		if err != nil {
			l.Error(":skull: Failed to get data source hash with error: %s", c, err.Error())
			return
		}

		rawRequests = append(rawRequests, rawRequest{
			dataSourceID:   raw.DataSourceID,
			dataSourceHash: hash,
			externalID:     raw.ExternalID,
			calldata:       string(raw.Calldata),
		})
	}

	// process raw requests
	reports, execVersions := handleRawRequests(c, l, id, rawRequests, key)

	offchainFeeUsed := rawRequests[0].offlineFeeLimit
	for idx := 1; idx < len(rawRequests); idx++ {
		for _, coin := range rawRequests[idx].offlineFeeLimit {
			for i, offchainCoin := range offchainFeeUsed {
				if offchainCoin.Denom == coin.Denom {
					offchainFeeUsed[i].Amount.Add(coin.Amount)
					break
				}
			}
		}
	}

	c.pendingMsgs <- ReportMsgWithKey{
		msg:         types.NewMsgReportData(types.RequestID(id), reports, c.validator, offchainFeeUsed),
		execVersion: execVersions,
		keyIndex:    keyIndex,
		feeEstimationData: FeeEstimationData{
			askCount:    int64(len(req.RequestedValidators)),
			minCount:    int64(req.MinCount),
			callData:    req.Calldata,
			rawRequests: rawRequests,
			clientID:    req.ClientID,
		},
	}
}

func handleRawRequests(
	c *Context,
	l *Logger,
	id types.RequestID,
	reqs []rawRequest,
	key keyring.Info,
) (reports []types.RawReport, execVersions []string) {
	resultsChan := make(chan processingResult, len(reqs))
	for _, req := range reqs {
		go handleRawRequest(
			c,
			l.With("did", req.dataSourceID, "eid", req.externalID),
			req,
			key,
			types.RequestID(id),
			resultsChan,
		)
	}

	versions := map[string]bool{}
	for range reqs {
		result := <-resultsChan
		reports = append(reports, result.rawReport)

		if result.err == nil {
			versions[result.version] = true
		}
	}

	for version := range versions {
		execVersions = append(execVersions, version)
	}

	return reports, execVersions
}

func handleRawRequest(
	c *Context,
	l *Logger,
	req rawRequest,
	key keyring.Info,
	id types.RequestID,
	processingResultCh chan processingResult,
) {
	c.updateHandlingGauge(1)
	defer c.updateHandlingGauge(-1)

	requirementFile, err := GetExecutable(c, l, req.requirementFileHash)
	if err != nil {
		l.Error(":skull: Failed to load requirement file with error: %s", c, err.Error())
		processingResultCh <- processingResult{
			rawReport: types.NewRawReport(
				req.externalID, 255, []byte("FAIL_TO_LOAD_REQUIREMENT_FILE"), nil, "", "",
			),
			err: err,
		}
		return

	}

	exec, err := GetExecutable(c, l, req.dataSourceHash)
	if err != nil {
		l.Error(":skull: Failed to load data source with error: %s", c, err.Error())
		processingResultCh <- processingResult{
			rawReport: types.NewRawReport(
				req.externalID, 255, []byte("FAIL_TO_LOAD_DATA_SOURCE"), nil, "", "",
			),
			err: err,
		}
		return
	}

	// send requirement file to runtime server to setup environment

	vmsg := types.NewRequestVerification(cfg.ChainID, c.validator, id, req.externalID, req.dataSourceID)
	sig, pubkey, err := kb.Sign(key.GetName(), vmsg.GetSignBytes())
	if err != nil {
		l.Error(":skull: Failed to sign verify message: %s", c, err.Error())
		processingResultCh <- processingResult{
			rawReport: types.NewRawReport(req.externalID, 255, nil, nil, "", ""),
			err:       err,
		}
		return
	}

	var executorObj executor.Executor
	for i, executorTemp := range c.executors {
		//c.mu.Lock()
		if executorTemp.GetLanguage() == req.language {
			c.executorIndex += 1
			executorObj = c.executors[c.executorIndex%len(c.executors)]
			break
		}

		if i == len(c.executors)-1 {
			l.Error(":skull: No runtime environment", c)
			processingResultCh <- processingResult{
				rawReport: types.NewRawReport(req.externalID, 255, nil, nil, "", ""),
				err:       err,
			}
			return
		}
	}

	executorObj.SetTimeout(req.baseOffchainFeePerHour, req.offlineFeeLimit.AmountOf("uband").Uint64())

	result, err := executorObj.Exec(req.usedExternalLibraries, req.baseOffchainFeePerHour, requirementFile, exec, req.calldata, map[string]interface{}{
		"BAND_CHAIN_ID":       vmsg.ChainID,
		"BAND_DATA_SOURCE_ID": strconv.Itoa(int(vmsg.DataSourceID)),
		"BAND_VALIDATOR":      vmsg.Validator,
		"BAND_REQUEST_ID":     strconv.Itoa(int(vmsg.RequestID)),
		"BAND_EXTERNAL_ID":    strconv.Itoa(int(vmsg.ExternalID)),
		"BAND_REPORTER":       hex.EncodeToString(pubkey.Bytes()),
		"BAND_SIGNATURE":      sig,
	})
	//c.mu.Unlock()

	if err != nil {
		l.Error(":skull: Failed to execute data source script: %s", c, err.Error())
		processingResultCh <- processingResult{
			rawReport: types.NewRawReport(req.externalID, 255, nil, nil, result.InstallationTime, result.ExecutionTime),
			err:       err,
		}
		return
	} else {
		l.Debug(
			":sparkles: Query data done with calldata: %q, result: %q, exitCode: %d",
			req.calldata, result.Output, result.Code,
		)

		offchainFeeUsed, err := sdk.ParseCoinsNormalized(result.OffchainFeeUsed)
		if err != nil {
			panic(err)
		}

		processingResultCh <- processingResult{
			rawReport: types.NewRawReport(req.externalID, result.Code, result.Output, offchainFeeUsed, result.InstallationTime, result.ExecutionTime),
			version:   result.Version,
		}
	}
}
