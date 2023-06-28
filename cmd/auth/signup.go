package auth

import (
	"fmt"
	"net/mail"
	"os"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/muesli/reflow/wrap"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/term"

	"github.com/onqlavelabs/onqlave.cli/cmd/common"
	"github.com/onqlavelabs/onqlave.cli/internal/api"
	"github.com/onqlavelabs/onqlave.cli/internal/utils"
	"github.com/onqlavelabs/onqlave.core/errors"
)

var (
	emailAddress string
	tenantName   string
	userFullName string
)

func validMailAddress(address string) bool {
	_, err := mail.ParseAddress(address)
	return err == nil
}

func signupCommand() *cobra.Command {
	init := &cobra.Command{
		Use:     "signup",
		Short:   "signup to platform by email",
		Long:    "This command is used to signup to platform. A valid email address, tenant name and full name are required. An invitation will be sent to the designated email.",
		Example: "onqlave auth signup",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return common.ReplacePersistentPreRunE(cmd, errors.NewCLIError(errors.KeyCLIMissingRequiredField, utils.BoldStyle.Render("Email address is required")))
			}
			if !validMailAddress(args[0]) {
				return common.ReplacePersistentPreRunE(cmd, errors.NewCLIError(errors.KeyCLIInvalidValue, utils.BoldStyle.Render("Email address is invalid. Please provide a valid email address")))
			}
			emailAddress = args[0]

			cmd.SilenceUsage = false

			return nil
		},
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			// Bind Cobra flags with viper
			if err := viper.BindPFlags(cmd.Flags()); err != nil {
				return common.ReplacePersistentPreRunE(cmd, err)
			}
			if !common.IsEnvironmentConfigured() {
				return common.ReplacePersistentPreRunE(cmd, common.ErrUnsetEnv)
			}
			if tenantName == "" {
				return common.ReplacePersistentPreRunE(cmd, errors.NewCLIError(errors.KeyCLIMissingRequiredField, utils.BoldStyle.Render("Tenant name should be provided")))
			}
			if userFullName == "" {
				return common.ReplacePersistentPreRunE(cmd, errors.NewCLIError(errors.KeyCLIMissingRequiredField, utils.BoldStyle.Render("User fullname should be provided")))
			}

			cmd.SilenceUsage = false

			return nil
		},
		Run: runSignupCommand,
	}
	init.Flags().StringVarP(&tenantName, "tenant_name", "t", "", "enter you tenant name, we will make a slug based on tenant name you provide to make it unique")
	init.Flags().StringVarP(&userFullName, "full_name", "n", "", "enter your full name")

	return init
}

func runSignupCommand(cmd *cobra.Command, args []string) {
	apiService := newAuthAPIService(cmd.Context())
	width, _, _ := term.GetSize(int(os.Stdout.Fd()))

	token, err := apiService.SendSignupInvitation(emailAddress, tenantName, userFullName)
	if err != nil {
		fmt.Println(utils.RenderError(fmt.Sprintf("There was an error sending the signup email to email address '%s': %s", emailAddress, err)) + "\n")
		return
	}

	s := &strings.Builder{}
	header := fmt.Sprintf("Signup instruction is sent to email address '%s'. Please be mindful that the link provided in email is only Valid for %d minutes.", emailAddress, common.Valid)
	s.WriteString(utils.BoldStyle.Copy().Foreground(utils.Color).Padding(1, 0, 0, 0).Render(wrap.String(header, width)))
	s.WriteString("\n")
	fmt.Println(s.String())

	communication := api.NewConcurrencyChannel()
	ui, err := utils.NewSpnnerTUI(cmd.Context(), utils.SpinnerOptions{
		Valid:    common.Valid,
		Consumer: communication.GetConsumer(),
	})
	if err != nil {
		fmt.Println(utils.RenderError(fmt.Sprintf("There was an error setting up signup operation: %s", err)) + "\n")
		return
	}

	go _waitingSignupOperation(apiService, token, communication.GetProducer(), common.Valid)

	if _, err := tea.NewProgram(ui).Run(); err != nil {

		fmt.Println(utils.RenderError(fmt.Sprintf("There was an error setting up signup operation: %s", err)) + "\n")
		return
	}

	if ui.Error() != nil {
		fmt.Println(utils.RenderError(fmt.Sprintf("There was an error whilst waiting for sign up result: %s", ui.Error())) + "\n")
	} else {
		fmt.Println(utils.BoldStyle.Copy().Foreground(utils.Green).Render("🎉 Done! You successfully signup to Onqlave platform. \n"))
	}

	fmt.Println(utils.TextStyle.Render("For more information, read our documentation at https://www.docs.onqlave.com \n"))
}

func _waitingSignupOperation(apiService *api.APIIntegrationService, token string, producer *api.Producer, valid int) {
	start := time.Now().UTC()
	duration := time.Since(start)
	producer.Produce(api.ConcurrencyOperationResult{Result: "Waiting for signup completion", Done: false, Error: nil})

	for duration.Minutes() < float64(valid) {
		result, err := apiService.GetSignupOperationStatus(token)
		producer.Produce(api.ConcurrencyOperationResult{Result: result.Result, Done: result.Done, Error: err})
		if result.Done || err != nil {
			return
		} else {
			time.Sleep(time.Second * 5)
		}
	}
}
