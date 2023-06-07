package arx

//
//import (
//	"errors"
//	"fmt"
//	"os"
//	"strings"
//	"time"
//
//	tea "github.com/charmbracelet/bubbletea"
//	"github.com/muesli/reflow/wrap"
//	"github.com/spf13/cobra"
//	"golang.org/x/term"
//
//	"github.com/onqlavelabs/onqlave.all/cmd/cli/commands/common"
//	"github.com/onqlavelabs/onqlave.all/internal/pkg/cli/api"
//	"github.com/onqlavelabs/onqlave.all/internal/pkg/cli/api/cluster"
//	"github.com/onqlavelabs/onqlave.all/internal/pkg/cli/cli"
//)
//
//type retryArxOperation struct {
//	arxId               string
//	arxOperationTimeout int
//}
//
//func (o retryArxOperation) waitForCompletion(apiService *cluster.ClusterAPIIntegrationService, arxId string, producer *api.Producer, valid int) {
//	start := time.Now().UTC()
//	duration := time.Since(start)
//	message := "Waiting for arx creation completion."
//	producer.Produce(api.ConcurrencyOperationResult{Result: message, Done: false, Error: nil})
//	for duration.Minutes() < float64(valid) {
//		result, err := apiService.CheckArxOperationState(arxId, cluster.RetryOperation)
//		producer.Produce(api.ConcurrencyOperationResult{Result: result.Result, Done: result.Done, Error: err})
//		if result.Done || err != nil {
//			return
//		} else {
//			time.Sleep(time.Second * 5)
//		}
//	}
//}
//
//var _retryArx retryArxOperation
//
//func retryCommand() *cobra.Command {
//	_retryArx.arxOperationTimeout = 10
//	return &cobra.Command{
//		Use:   "retry",
//		Short: "retry adding arx by ID and attributes",
//		Long: "This command is used to retry adding arx by ID. Valid arx id, arx name, arx provider,  arx type, arx purpose, " +
//			"arx region, arx description, arx encryption method, arx rotation cycle, arx owner, arx spend limit " +
//			"and arx is default are required.",
//		Example: "onqlave arx retry",
//		Args: func(cmd *cobra.Command, args []string) error {
//			if len(args) < 1 {
//				cmd.SilenceUsage = true
//
//				return errors.New("arx id is required")
//			}
//			_retryArx.arxId = args[0]
//			return nil
//		},
//		Run: runRetryCommand,
//	}
//}
//
//func runRetryCommand(cmd *cobra.Command, args []string) {
//	width, _, _ := term.GetSize(int(os.Stdout.Fd()))
//	arxID := _retryArx.arxId
//
//	apiService := cluster.NewClusterAPIIntegrationService(cluster.ClusterAPIIntegrationServiceOptions{Ctx: cmd.Context()})
//
//	arxId, err := apiService.RetryAddCluster(arxID)
//	if err != nil {
//		common.RenderCLIOutputError(fmt.Sprintf("There was an error retry creating arx '%s': ", arxID), err)
//		return
//	}
//	s := &strings.Builder{}
//	header := fmt.Sprintf("Arx creation sometime takes up to %d minutes.", _retryArx.arxOperationTimeout)
//	s.WriteString(cli.BoldStyle.Copy().Foreground(cli.Color).Padding(1, 0, 0, 0).Render(wrap.String(header, width)))
//	fmt.Println(s.String())
//
//	communication := api.NewConcurrencyChannel()
//	// Run the function.
//	ui, err := cli.NewSpnnerTUI(cmd.Context(), cli.SpinnerOptions{
//		Valid:    common.Valid,
//		Consumer: communication.GetConsumer(),
//	})
//	if err != nil {
//		fmt.Println(cli.RenderError(fmt.Sprintf("There was an error setting up arx creation operation: %s", err)) + "\n")
//		return
//	}
//	go func() {
//		_retryArx.waitForCompletion(apiService, arxId, communication.GetProducer(), _retryArx.arxOperationTimeout)
//	}()
//	if _, err := tea.NewProgram(ui).Run(); err != nil {
//		fmt.Println(cli.RenderError(fmt.Sprintf("There was an error setting up arx creation operation: %s", err)) + "\n")
//		return
//	}
//
//	common.CliRenderUIErrorOutput(ui, common.ResourceArx, common.ActionCreated, arxID)
//}
