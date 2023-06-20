package common

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/onqlavelabs/onqlave.cli/core/contracts/common"
	coreErr "github.com/onqlavelabs/onqlave.cli/core/errors"
	"github.com/onqlavelabs/onqlave.cli/internal/pkg/cli/cli"
	"github.com/onqlavelabs/onqlave.cli/internal/pkg/model"
)

var (
	ErrUnsupportedEnv = errors.New("environment is invalid. It should be either 'dev' or 'prod'")
	UnsetEnvError     = errors.New(cli.BoldStyle.Render(`your environment is not configured | you are not logged in to the environment. Please run 'config init | config auth' before running any other command`))
)

func GetStatusAndMessageErr(jsonString string) (string, string, error) {
	jsonBytes := []byte(jsonString)
	var jsonObj common.BaseErrorResponse

	err := json.Unmarshal(jsonBytes, &jsonObj)
	if err != nil {
		return "", "", errors.New(jsonString)
	}

	status := jsonObj.Error.Status
	message := jsonObj.Error.Message
	return status, message, nil
}

func RenderCLIOutputError(prefix string, err error) {
	appErr, ok := err.(*model.AppError)
	if !ok {
		fmt.Println(cli.RenderError(fmt.Sprintf("%s", err)) + "\n")
		return
	}

	var result string
	unwrapError := appErr.Unwrap()
	status, message, err := GetStatusAndMessageErr(fmt.Sprintf("%s", unwrapError))
	if err != nil {
		fmt.Println(cli.RenderError(fmt.Sprintf("%s%s", prefix, err)) + "\n")
		return
	}

	result = fmt.Sprintf("%s%s", prefix, message)

	if status == coreErr.KeyServiceDecryptErr {
		result = fmt.Sprintf("%s%s", prefix, "You are unauthorized to perform this action")
	}

	fmt.Println(cli.RenderError(result) + "\n")
}
