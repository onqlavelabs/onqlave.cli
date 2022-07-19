package commands

import (
	"errors"
	"fmt"
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

type retryClusterOperation struct {
	clusterId               string
	clusterOperationTimeout int
}

func (o retryClusterOperation) waitForCompletion(apiService *api.APIIntegrationService, clusterId string, producer *api.Producer, valid int) {
	start := time.Now()
	duration := time.Since(start)
	message := "Waiting for cluster creation completion."
	producer.Produce(api.ConcurrencyOperationResult{Result: message, Done: false, Error: nil})
	for duration.Minutes() < float64(valid) {
		result, err := apiService.AddClusterOperationState(clusterId)
		producer.Produce(api.ConcurrencyOperationResult{Result: result.Result, Done: result.Done, Error: err})
		if result.Done || err != nil {
			return
		} else {
			time.Sleep(time.Second * 5)
		}
	}
}

var _retryCluster retryClusterOperation

func clusterRetryCommand() *cobra.Command {
	_retryCluster.clusterOperationTimeout = 10
	init := &cobra.Command{
		Use:     "retry",
		Short:   "retry.",
		Long:    "retry command.",
		Example: "onqlave clusters retry",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("clusterid is required")
			}
			_retryCluster.clusterId = args[0]
			return nil
		},
		// used to overwrite/skip the parent commands persistentPreRunE func
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			// Bind Cobra flags with viper
			if err := viper.BindPFlags(cmd.Flags()); err != nil {
				return err
			}
			if !isLoggedin() {
				return errors.New("your environment is not configured | you are not logged in to the environment. please run 'config init | config auth' before running any other command")
			}
			return nil
		},
		Run: runClusterRetryCommand,
	}
	return init
}

func runClusterRetryCommand(cmd *cobra.Command, args []string) {
	width, _, _ := term.GetSize(int(os.Stdout.Fd()))

	apiService := api.NewAPIIntegrationService(api.APIIntegrationServiceOptions{Ctx: cmd.Context()})

	clusterId, err := apiService.RetryAddCluster(_retryCluster.clusterId)
	if err != nil {
		fmt.Println(cli.RenderError(fmt.Sprintf("There was an error retry creating cluster '%s': %s", _retryCluster.clusterId, err)) + "\n")
		return
	}
	s := &strings.Builder{}
	header := fmt.Sprintf("Cluster creation sometime takes up to %d minutes.", _retryCluster.clusterOperationTimeout)
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
		fmt.Println(cli.RenderError(fmt.Sprintf("There was an error setting up cluster creation operation: %s", err)) + "\n")
		return
	}
	go func() {
		_retryCluster.waitForCompletion(apiService, clusterId, communication.GetProducer(), _retryCluster.clusterOperationTimeout)
	}()
	if err := tea.NewProgram(ui).Start(); err != nil {
		fmt.Println(cli.RenderError(fmt.Sprintf("There was an error setting up cluster creation operation: %s", err)) + "\n")
		return
	}
	if ui.Error() != nil {
		fmt.Println(cli.RenderError(fmt.Sprintf("There was an error whilst waiting for cluster creation result: %s", ui.Error())) + "\n")
	} else {
		fmt.Println(cli.BoldStyle.Copy().Foreground(cli.Green).Render("🎉 Done!  Cluster created successfully. \n"))
	}
	fmt.Println(cli.TextStyle.Render("For more information, read our documentation at https://www.onqlave.com/docs\n"))
}
