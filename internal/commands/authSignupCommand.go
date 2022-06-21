package commands

import (
	"errors"
	"fmt"
	"net/mail"
	"os"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/muesli/reflow/wrap"
	"github.com/onqlavelabs/onqlave.core/internal/api"
	"github.com/onqlavelabs/onqlave.core/internal/cli"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/term"
)

var (
	emailAddress string
	tenantName   string
)

func validMailAddress(address string) bool {
	_, err := mail.ParseAddress(address)
	return err == nil
}

func signupCommand() *cobra.Command {
	init := &cobra.Command{
		Use:     "signup",
		Short:   "signup.",
		Long:    "signup command.",
		Example: "onqlave auth signup",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("requires email address")
			}
			if !validMailAddress(args[0]) {
				return errors.New("email address is invalid. Please provide a valid email address")
			}
			emailAddress = args[0]
			return nil
		},
		// used to overwrite/skip the parent commands persistentPreRunE func
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			// Bind Cobra flags with viper
			if err := viper.BindPFlags(cmd.Flags()); err != nil {
				return err
			}
			if tenantName == "" {
				return fmt.Errorf("tenant name should be provided")
			}
			return nil
		},
		Run: runConfigSignupCommand,
	}
	// init.Flags().StringVarP(&emailAddress, "email-address", "e", "", "Enter your email address to signup")
	init.Flags().StringVarP(&tenantName, "tenant_name", "t", "", "Enter you tenant name. We will make a slug based on tenant name you provide to make it unqiue")
	return init
}

func runConfigSignupCommand(cmd *cobra.Command, args []string) {
	width, _, _ := term.GetSize(int(os.Stdout.Fd()))

	apiService := api.NewAPIIntegrationService(api.APIIntegrationServiceOptions{Ctx: cmd.Context()})

	token, err := apiService.SendSignupInvitation(emailAddress, tenantName)
	if err != nil {
		fmt.Println(cli.RenderError(fmt.Sprintf("There was an error sending the signup email to email address '%s': %s", emailAddress, err)) + "\n")
		return
	}

	s := &strings.Builder{}
	header := fmt.Sprintf("Signup instruction is sent to email address '%s'. Please be mindful that the link provided in email is only valid for %d minutes.", emailAddress, valid)
	s.WriteString(cli.BoldStyle.Copy().Foreground(cli.Color).Padding(1, 0, 0, 0).Render(wrap.String(header, width)))
	s.WriteString("\n")
	fmt.Println(s.String())

	communication := api.NewConcurrencyChannel()
	// Run the function.
	ui, err := cli.NewSpnnerTUI(cmd.Context(), cli.SpinnerOptions{
		Valid:    valid,
		Consumer: communication.GetConsumer(),
	})
	if err != nil {
		fmt.Println(cli.RenderError(fmt.Sprintf("There was an error setting up signup operation: %s", err)) + "\n")
		return
	}
	go func() {
		_waitingSignupOperation(apiService, token, communication.GetProducer(), valid)
	}()
	if err := tea.NewProgram(ui).Start(); err != nil {
		fmt.Println(cli.RenderError(fmt.Sprintf("There was an error setting up signup operation: %s", err)) + "\n")
		return
	}
	if ui.Error() != nil {
		fmt.Println(cli.RenderError(fmt.Sprintf("There was an error whilst waiting for sign up result: %s", ui.Error())) + "\n")
	} else {
		fmt.Println(cli.BoldStyle.Copy().Foreground(cli.Green).Render("🎉 Done!  You successfully signup to Onqlave platform. \n"))
	}
	fmt.Println(cli.TextStyle.Render("For more information, read our documentation at https://www.onqlave.com/docs\n"))
}

func _waitingSignupOperation(apiService *api.APIIntegrationService, token string, producer *api.Producer, valid int) {
	start := time.Now()
	duration := time.Since(start)
	producer.Produce(api.ConcurrencyOperationResult{Result: "Waiting for signup completion", Done: false, Error: nil})
	for duration.Minutes() < float64(valid) {
		result, err := apiService.GetSignupOperationStatus(token)
		time.Sleep(time.Second * 5)
		producer.Produce(api.ConcurrencyOperationResult{Result: result.Result, Done: result.Done, Error: err})
		if result.Done {
			return
		}
	}
}
