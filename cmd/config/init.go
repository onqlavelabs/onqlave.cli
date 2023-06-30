package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/onqlavelabs/onqlave.cli/cmd/common"
	"github.com/onqlavelabs/onqlave.cli/internal/config"
	"github.com/onqlavelabs/onqlave.cli/internal/utils"
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
	case common.EnvLocal:
		baseUrl = common.BaseUrlLocal
	default:
		cliEnv = common.EnvProd
		baseUrl = common.BaseUrlProd
	}

	viper.Set(common.FlagEnv, cliEnv)
	viper.Set(common.FlagApiBaseUrl, baseUrl)

	err := config.CreateFile(viper.GetString(common.FlagConfigPath))
	if err != nil {
		return common.ReplacePersistentPreRunE(cmd, err)
	}

	fmt.Println("")
	fmt.Println(utils.BoldStyle.Copy().Foreground(utils.Green).Render("🎉 Done! You successfully initialize your environment . Next step is to signup or login.\n"))
	fmt.Println(utils.TextStyle.Render("For more information, read our documentation at https://docs.onqlave.com \n"))

	return nil
}
