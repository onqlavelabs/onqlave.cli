package auth

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/muesli/reflow/wrap"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/term"

	"github.com/onqlavelabs/onqlave.cli/cmd/common"
	"github.com/onqlavelabs/onqlave.cli/internal/pkg/cli/api"
	"github.com/onqlavelabs/onqlave.cli/internal/pkg/cli/cli"
	"github.com/onqlavelabs/onqlave.cli/internal/pkg/cli/configs"
)

func loginCommand() *cobra.Command {
	init := &cobra.Command{
		Use:     "login",
		Short:   "login to platform by email",
		Long:    "This command is used to login to platform. A valid email address, tenant name are required. An invitation will be sent to the designated email.",
		Example: "onqlave auth login",
		Args: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true

			if len(args) < 1 {
				return errors.New("requires email address")
			}
			if !validMailAddress(args[0]) {
				return errors.New("email address is invalid. Please provide a valid email address")
			}
			emailAddress = args[0]

			cmd.SilenceUsage = false

			return nil
		},
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true

			if err := viper.BindPFlags(cmd.Flags()); err != nil {
				return common.ReplacePersistentPreRunE(cmd, err)
			}
			if !common.IsEnvironmentConfigured() {
				return common.ReplacePersistentPreRunE(cmd, errors.New("your environment is not configured. please run 'config init' before running any other command"))
			}
			if tenantName == "" {
				return common.ReplacePersistentPreRunE(cmd, errors.New("tenant name should be provided"))

			}

			cmd.SilenceUsage = false

			return nil
		},
		Run: runLoginCommand,
	}

	init.Flags().StringVarP(&tenantName, "tenant_name", "t", "", "enter you tenant name, we will make a slug based on tenant name you provide to make it unique")

	return init
}

func runLoginCommand(cmd *cobra.Command, args []string) {
	apiService := newAuthAPIService(cmd.Context())
	width, _, _ := term.GetSize(int(os.Stdout.Fd()))

	token, err := apiService.SendLoginInvitation(emailAddress, tenantName)
	if err != nil {
		fmt.Println(cli.RenderError(fmt.Sprintf("There was an error sending the login email to email address '%s': %s", emailAddress, err)) + "\n")
		return
	}

	s := &strings.Builder{}
	header := fmt.Sprintf("Login instruction is sent to email address '%s'. Please be mindful that the link provided in email is only valid for %d minutes.", emailAddress, common.Valid)
	s.WriteString(cli.BoldStyle.Copy().Foreground(cli.Color).Padding(1, 0, 0, 0).Render(wrap.String(header, width)))
	s.WriteString("\n")
	fmt.Println(s.String())

	communication := api.NewConcurrencyChannel()
	ui, err := cli.NewSpnnerTUI(cmd.Context(), cli.SpinnerOptions{Valid: common.Valid, Consumer: communication.GetConsumer()})
	if err != nil {
		fmt.Println(cli.RenderError(fmt.Sprintf("There was an error setting up login operation: %s", err)) + "\n")
		return
	}

	go _waitingLoginOperation(apiService, token, communication.GetProducer(), common.Valid)

	if _, err := tea.NewProgram(ui).Run(); err != nil {
		fmt.Println(cli.RenderError(fmt.Sprintf("There was an error setting up login operation: %s", err)) + "\n")
		return
	}

	if ui.Error() != nil {
		fmt.Println(cli.RenderError(fmt.Sprintf("There was an error whilst waiting for login result: %s", ui.Error())) + "\n")
	} else {
		fmt.Println(cli.BoldStyle.Copy().Foreground(cli.Green).Render("🎉 Done! You successfully login to Onqlave platform. \n"))
	}

	fmt.Println(cli.TextStyle.Render("For more information, read our documentation at https://docs.onqlave.com \n"))
}

func _waitingLoginOperation(apiService *api.APIIntegrationService, token string, producer *api.Producer, valid int) {
	start := time.Now().UTC()
	duration := time.Since(start)
	producer.Produce(api.ConcurrencyOperationResult{Result: "Waiting for login completion", Done: false, Error: nil})

	for duration.Minutes() < float64(valid) {
		result, authToken, tenantID, err := apiService.GetLoginOperationStatus(token)
		producer.Produce(api.ConcurrencyOperationResult{Result: result.Result, Done: result.Done, Error: err})
		if result.Done || err != nil {
			if authToken != "" && err == nil {
				viper.Set(common.FlagAuthKey, authToken)
				viper.Set(common.FlagTenantID, tenantID)
				err := configs.CreateFile(viper.GetString(common.FlagConfigPath))
				if err != nil {
					return
				}
			}
			return
		} else {
			time.Sleep(time.Second * 5)
		}
	}

}