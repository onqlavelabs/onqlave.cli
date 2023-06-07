package config

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/muesli/reflow/wrap"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/term"

	cliCommon "github.com/onqlavelabs/onqlave.cli/cmd/common"
	"github.com/onqlavelabs/onqlave.cli/internal/pkg/cli/cli"
)

type CurrentConfig struct {
	ApiBaseUrl string `json:"api_base_url"`
	ConfigPath string `json:"config_path"`
	Env        string `json:"env"`
	TenantId   string `json:"tenant_id"`
	TenantName string `json:"tenant_name"`
}

func currentCommand() *cobra.Command {
	currentCmd := &cobra.Command{
		Use:     "current",
		Short:   "get your current environment configuration",
		Long:    "This command is used to get the current environment configuration",
		Example: "onqlave config current",
		Run:     runCurrentCommand,
	}

	return currentCmd
}

func runCurrentCommand(cmd *cobra.Command, args []string) {
	width, _, _ := term.GetSize(int(os.Stdout.Fd()))

	configInfo := getConfigInfo()

	if viper.GetBool(cliCommon.FlagJson) {
		s := &strings.Builder{}
		s.WriteString(cli.BoldStyle.Copy().Foreground(cli.Color).Padding(1, 0, 0, 0).Render(wrap.String("Current Config information =>", width)))
		s.WriteString("\n")
		s.WriteString(cli.RenderAsJson(configInfo))
		s.WriteString("\n")
		fmt.Println(s.String())
		return
	}

	cliCommon.NewDataTable(configInfo).Render()
}

func getConfigInfo() CurrentConfig {
	url := viper.Get(cliCommon.FlagApiBaseUrl)
	apiBaseUrl := viper.GetString(cliCommon.FlagApiBaseUrl)
	if reflect.TypeOf(url).Kind() == reflect.Map {
		var m struct {
			Billing       string `json:"billings"`
			Registrations string `json:"registrations"`
			Tenants       string `json:"tenants"`
		}
		byteApiBaseUrl, _ := json.Marshal(url)
		_ = json.Unmarshal(byteApiBaseUrl, &m)
		apiBaseUrl = fmt.Sprintf("%v", m)
	}

	return CurrentConfig{
		ApiBaseUrl: apiBaseUrl,
		Env:        viper.GetString(cliCommon.FlagEnv),
		ConfigPath: viper.GetString(cliCommon.FlagConfigPath),
		TenantId:   viper.GetString(cliCommon.FlagTenantID),
		TenantName: viper.GetString("tenant_name"),
	}
}
