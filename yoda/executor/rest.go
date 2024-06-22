package executor

import (
	"encoding/base64"
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/levigross/grequests"
)

type RestExec struct {
	url     string
	timeout time.Duration
}

func NewRestExec(url string, timeoutFromYoda time.Duration) *RestExec {
	//timeout := timeoutFromRequest
	//if timeoutFromYoda < timeoutFromRequest {
	//	timeout = timeoutFromYoda
	//}
	return &RestExec{url: url, timeout: timeoutFromYoda}
}

type externalExecutionResponse struct {
	Returncode uint32 `json:"returncode"`
	Stdout     string `json:"stdout"`
	Stderr     string `json:"stderr"`
	Duration   string `json:"duration"`
	Version    string `json:"version"`
	Error      string `json:"err"`
}

func (e *RestExec) SetTimeout(baseOffchainFeePerHour uint64, amt uint64) {
	timeoutFromRequest := time.Duration(amt / baseOffchainFeePerHour * uint64(time.Hour))
	if e.timeout > timeoutFromRequest {
		e.timeout = timeoutFromRequest
	}
}

func (e *RestExec) Exec(baseOffchainFeePerHour uint64, requirementFile []byte, code []byte, arg string, env interface{}) (ExecResult, error) {
	executable := base64.StdEncoding.EncodeToString(code)
	resp, err := grequests.Post(
		e.url,
		&grequests.RequestOptions{
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
			JSON: map[string]interface{}{
				"requirement-file": requirementFile,
				"executable":       executable,
				"calldata":         arg,
				"timeout":          e.timeout.Milliseconds(),
				"env":              env,
			},
			RequestTimeout: e.timeout,
		},
	)

	fmt.Println(resp)

	if err != nil {
		urlErr, ok := err.(*url.Error)
		if !ok || !urlErr.Timeout() {
			return ExecResult{}, err
		}
		// Return timeout code
		offchainFeeUsed := e.timeout.Hours() * float64(baseOffchainFeePerHour)
		offchainFeeUsedStr := fmt.Sprintf("%suband", strconv.FormatUint(uint64(offchainFeeUsed), 10))
		return ExecResult{Output: []byte{}, Code: 111, OffchainFeeUsed: offchainFeeUsedStr}, nil
	}

	if !resp.Ok {
		return ExecResult{}, ErrRestNotOk
	}

	r := externalExecutionResponse{}
	err = resp.JSON(&r)

	if err != nil {
		return ExecResult{}, err
	}

	duration, err := time.ParseDuration(r.Duration)
	if err != nil {
		return ExecResult{}, err
	}
	offchainFeeUsed := duration.Hours() * float64(baseOffchainFeePerHour)
	offchainFeeUsedStr := fmt.Sprintf("%suband", strconv.FormatUint(uint64(offchainFeeUsed), 10))

	if r.Returncode == 0 {
		return ExecResult{Output: []byte(r.Stdout), OffchainFeeUsed: offchainFeeUsedStr, Code: 0, Version: r.Version}, nil
	} else {
		return ExecResult{Output: []byte(r.Stderr), OffchainFeeUsed: offchainFeeUsedStr, Code: r.Returncode, Version: r.Version}, nil
	}
}
