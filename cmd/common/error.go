package common

import (
	"encoding/json"
	"fmt"

	"github.com/onqlavelabs/onqlave.cli/internal/model"
	"github.com/onqlavelabs/onqlave.cli/internal/utils"
	"github.com/onqlavelabs/onqlave.core/contracts/common"
	"github.com/onqlavelabs/onqlave.core/errors"
)

var (
	ErrUnsupportedEnv = errors.NewCLIError(errors.KeyCLIInvalidValue, utils.BoldStyle.Render("Environment is invalid. It should be either 'dev' or 'prod'"))
	ErrUnsetEnv       = errors.NewCLIError(errors.KeyCLIEnvironmentNotConfig, utils.BoldStyle.Render(`Your environment is not configured. Please run 'config init' before running any other command`))
	ErrRequireLogIn   = errors.NewCLIError(errors.KeyCLINotLoggedIn, utils.BoldStyle.Render(`You are not logged in to the environment. Please run 'auth login' before running any other command`))
)

func GetStatusAndMessageErr(jsonString string) (string, string, error) {
	jsonBytes := []byte(jsonString)
	var jsonObj common.BaseErrorResponse

	err := json.Unmarshal(jsonBytes, &jsonObj)
	if err != nil {
		return "", "", errors.NewCLIError(errors.KeyServiceErr, jsonString)
	}

	status := jsonObj.Error.Status
	message := jsonObj.Error.Message
	return status, message, nil
}

func RenderCLIOutputError(prefix string, err error) {
	appErr, ok := err.(*model.AppError)
	if !ok {
		fmt.Println(utils.RenderError(fmt.Sprintf("%s", err)) + "\n")
		return
	}

	var result string
	unwrapError := appErr.Unwrap()
	status, message, err := GetStatusAndMessageErr(fmt.Sprintf("%s", unwrapError))
	if err != nil {
		fmt.Println(utils.RenderError(fmt.Sprintf("%s%s", prefix, err)) + "\n")
		return
	}

	result = fmt.Sprintf("%s%s", prefix, message)

	if status == errors.KeyServiceDecryptErr {
		result = fmt.Sprintf("%s%s", prefix, "You are unauthorized to perform this action")
	}

	fmt.Println(utils.RenderError(result) + "\n")
}
