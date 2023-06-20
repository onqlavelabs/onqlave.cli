package config

import (
	"fmt"
	"github.com/onqlavelabs/onqlave.cli/internal/cli/cli"
	"github.com/onqlavelabs/onqlave.cli/internal/cli/configs"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/onqlavelabs/onqlave.cli/cmd/common"
)

func initCommand() *cobra.Command {
	initCmd := &cobra.Command{
		Use:     "init",
		Short:   "initialize environment variables",
		Long:    "This command is used to initialize your CLI environment variables.",
		Example: "onqlave config init",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			if err := viper.BindPFlags(cmd.Flags()); err != nil {
				return common.ReplacePersistentPreRunE(cmd, err)
			}

			if common.IsSupportedEnv() {
				return common.ReplacePersistentPreRunE(cmd, common.ErrUnsupportedEnv)
			}

			cmd.SilenceUsage = false

			return nil
		},
		RunE: runInitCommand,
	}

	initCmd.PersistentFlags().StringP(common.FlagConfigPath, "c", common.GetConfigFilePath(), "location of the onqlave config file")

	return initCmd
}

func runInitCommand(cmd *cobra.Command, args []string) error {
	var baseUrl interface{}
	cliEnv := os.Getenv(strings.ToUpper(common.FlagEnv))

	switch cliEnv {
	case common.EnvDev:
		baseUrl = common.BaseUrlDev
	case common.EnvProd:
		baseUrl = common.BaseUrlProd
	default:
		baseUrl = common.BaseUrlLocal
		cliEnv = common.EnvLocal
	}

	viper.Set(common.FlagEnv, cliEnv)
	viper.Set(common.FlagApiBaseUrl, baseUrl)

	err := configs.CreateFile(viper.GetString(common.FlagConfigPath))
	if err != nil {
		return common.ReplacePersistentPreRunE(cmd, err)
	}

	fmt.Println("")
	fmt.Println(cli.BoldStyle.Copy().Foreground(cli.Green).Render("🎉 Done! You successfully initialize your environment . Next step is to signup/login is you already haven't.\n"))
	fmt.Println(cli.TextStyle.Render("For more information, read our documentation at https://docs.onqlave.com \n"))

	return nil
}
