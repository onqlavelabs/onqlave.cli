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

type addClusterOperation struct {
	clusterName             string
	clusterType             string
	clusterProvider         string
	clusterPurpose          string
	clusterRegion           string
	clusterDescription      string
	clusterOperationTimeout int
}

func (o addClusterOperation) waitForCompletion(apiService *api.APIIntegrationService, clusterId string, producer *api.Producer, valid int) {
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

var _addCluster addClusterOperation

func clusterAddCommand() *cobra.Command {
	_addCluster.clusterOperationTimeout = 10
	init := &cobra.Command{
		Use:     "add",
		Short:   "add.",
		Long:    "add command.",
		Example: "onqlave clusters add",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("clustername is required")
			}
			_addCluster.clusterName = args[0]
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
			if api.NewClusterProvider(_addCluster.clusterProvider) == api.ProviderInvalid {
				return errors.New("invalid cluster provider - must be in (GCP, AWS, AZURE)")
			}

			if api.NewClusterType(_addCluster.clusterType) == api.InvalidCluster {
				return errors.New("invalid cluster type - must be in (Serverless)")
			}

			if api.NewClusterRegion(_addCluster.clusterRegion) == api.REGION_INVALID {
				return errors.New("invalid cluster region - must be in (AUS-EAST, AUS-WEST)")
			}

			if api.NewClusterPurpose(_addCluster.clusterPurpose) == api.PurposeInvalid {
				return errors.New("invalid cluster purpose - must be in (Testing)")
			}
			return nil
		},
		Run: runClusterAddCommand,
	}
	init.Flags().StringVarP(&_addCluster.clusterProvider, "cluster_provider", "p", "GCP", "Enter Cluster Cloud Provider - (GCP, AWS, AZURE)")
	init.Flags().StringVarP(&_addCluster.clusterType, "cluster_type", "t", "Serverless", "Enter Cluster Type - (Serverless)")
	init.Flags().StringVarP(&_addCluster.clusterPurpose, "cluster_purpose", "u", "Testing", "Enter Cluster Purpose - (Testing)")
	init.Flags().StringVarP(&_addCluster.clusterRegion, "cluster_region", "r", "", "Enter Cluster Region - (AUS-EAST, AUS-WEST)")
	init.Flags().StringVarP(&_addCluster.clusterDescription, "cluster_description", "d", "Default", "Enter Cluster Description")

	return init
}

func runClusterAddCommand(cmd *cobra.Command, args []string) {
	width, _, _ := term.GetSize(int(os.Stdout.Fd()))

	apiService := api.NewAPIIntegrationService(api.APIIntegrationServiceOptions{Ctx: cmd.Context()})

	clusterId, _, err := apiService.AddCluster(_addCluster.clusterName,
		_addCluster.clusterProvider,
		_addCluster.clusterType,
		_addCluster.clusterPurpose,
		_addCluster.clusterRegion,
		_addCluster.clusterDescription)
	if err != nil {
		fmt.Println(cli.RenderError(fmt.Sprintf("There was an error creating cluster '%s': %s", _addCluster.clusterName, err)) + "\n")
		return
	}
	// start := time.Now()
	// duration := time.Since(start)
	// for duration.Minutes() < float64(valid) {
	// 	result, err := apiService.AddClusterOperationState(clusterId)
	// 	if result.Done || err != nil {
	// 		fmt.Printf("error %s", err)
	// 		return
	// 	} else {
	// 		time.Sleep(time.Second * 5)
	// 	}
	// }

	s := &strings.Builder{}
	header := fmt.Sprintf("Cluster creation sometime takes up to %d minutes.", _addCluster.clusterOperationTimeout)
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
		_addCluster.waitForCompletion(apiService, clusterId, communication.GetProducer(), _addCluster.clusterOperationTimeout)
	}()
	if err := tea.NewProgram(ui).Start(); err != nil {
		fmt.Println(cli.RenderError(fmt.Sprintf("There was an error setting up cluster creation operation: %s", err)) + "\n")
		return
	}
	if ui.Error() != nil {
		fmt.Println(cli.RenderError(fmt.Sprintf("There was an error whilst waiting for cluster creation result: %s", ui.Error())) + "\n")
		fmt.Println(cli.TextStyle.Render(fmt.Sprintf("You can re-initiate the cluster creation workflow by running 'onqlave clusters retry %s' \n", clusterId)))
	} else {
		fmt.Println(cli.BoldStyle.Copy().Foreground(cli.Green).Render("🎉 Done!  Cluster created successfully. \n"))
	}
	fmt.Println(cli.TextStyle.Render("For more information, read our documentation at https://www.onqlave.com/docs\n"))
}
