package executor

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
	"time"
)

const (
	flagQueryTimeout  = "timeout"
	flagQueryLanguage = "lang"
)

var (
	ErrExecutionimeout = errors.New("execution timeout")
	ErrRestNotOk       = errors.New("rest return non 2XX response")
)

type ExecResult struct {
	Output           []byte
	Code             uint32
	OffchainFeeUsed  string
	Version          string
	InstallationTime string
	ExecutionTime    string
}

type Executor interface {
	Exec(usedExternalLibraries string, baseOffchainFeePerHour uint64, requirementFile []byte, exec []byte, arg string, env interface{}) (ExecResult, error)
	SetTimeout(baseOffchainFeePerHour uint64, amt uint64)
	GetLanguage() string
}

var testProgram []byte = []byte(
	"import os\nimport sys\nprint(sys.argv[1], os.getenv('BAND_CHAIN_ID'))",
)

var testRequirementFile []byte = []byte(
	"requests==2.28.0\nurllib3==1.26.9\nwebsocket-client==1.3.2\nyarl==1.9.2",
)

// NewExecutor returns executor by name and executor URL
func NewExecutors(executorsStr string) ([]Executor, error) {
	execs := make([]Executor, 0)
	executors := strings.Split(executorsStr, ",")

	for _, executor := range executors {
		name, base, timeout, language, err := parseExecutor(executor)
		if err != nil {
			return nil, err
		}
		switch name {
		case "rest":
			exec := NewRestExec(base, timeout, language)
			execs = append(execs, exec)
		case "docker":
			return nil, fmt.Errorf("Docker executor is currently not supported")
		default:
			return nil, fmt.Errorf("Invalid executor name: %s, base: %s", name, base)
		}
	}

	// TODO: Remove hardcode in test execution
	//res, err := exec.Exec(DefaultFeePerHour, testRequirementFile, testProgram, "TEST_ARG", map[string]interface{}{
	//	"BAND_CHAIN_ID":    "test-chain-id",
	//	"BAND_VALIDATOR":   "test-validator",
	//	"BAND_REQUEST_ID":  "test-request-id",
	//	"BAND_EXTERNAL_ID": "test-external-id",
	//	"BAND_REPORTER":    "test-reporter",
	//	"BAND_SIGNATURE":   "test-signature",
	//})
	//
	//if err != nil {
	//	return nil, fmt.Errorf("failed to run test program: %s", err.Error())
	//}
	//if res.Code != 0 {
	//	return nil, fmt.Errorf("test program returned nonzero code: %d", res.Code)
	//}
	//if string(res.Output) != "TEST_ARG test-chain-id\n" {
	//	return nil, fmt.Errorf("test program returned wrong output: {%s}", res.Output)
	//}
	return execs, nil
}

// parseExecutor splits the executor string in the form of "name:base?timeout=" into parts.
func parseExecutor(executorStr string) (name string, base string, timeout time.Duration, language string, err error) {
	executor := strings.SplitN(executorStr, ":", 2)
	if len(executor) != 2 {
		return "", "", 0, "", fmt.Errorf("Invalid executor, cannot parse executor: %s", executorStr)
	}
	u, err := url.Parse(executor[1])
	if err != nil {
		return "", "", 0, "", fmt.Errorf("Invalid url, cannot parse %s to url with error: %s", executor[1], err.Error())
	}

	query := u.Query()
	timeoutStr := query.Get(flagQueryTimeout)
	if timeoutStr == "" {
		return "", "", 0, "", fmt.Errorf("Invalid timeout, executor requires query timeout")
	}

	language = query.Get(flagQueryLanguage)
	if language == "" {
		language = "python"
	}
	// Remove timeout from query because we need to return `base`
	query.Del(flagQueryTimeout)
	query.Del(flagQueryLanguage)
	u.RawQuery = query.Encode()

	timeout, err = time.ParseDuration(timeoutStr)
	if err != nil {
		return "", "", 0, "", fmt.Errorf("Invalid timeout, cannot parse duration with error: %s", err.Error())
	}
	return executor[0], u.String(), timeout, language, nil
}
