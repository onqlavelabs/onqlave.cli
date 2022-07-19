package commands

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/muesli/reflow/wrap"
	"github.com/onqlavelabs/onqlave.core/internal/api"
	"github.com/onqlavelabs/onqlave.core/internal/cli"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/term"
)

type describeClusterOperation struct {
	clusterId string
}

var _describeCluster describeClusterOperation

func clusterDescribeCommand() *cobra.Command {
	init := &cobra.Command{
		Use:     "describe",
		Short:   "describe.",
		Long:    "describe command.",
		Example: "onqlave cluster describe",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("clusterid is required")
			}
			_describeCluster.clusterId = args[0]
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
		Run: runClusterDescribeCommand,
	}
	return init
}

func runClusterDescribeCommand(cmd *cobra.Command, args []string) {
	width, _, _ := term.GetSize(int(os.Stdout.Fd()))

	apiService := api.NewAPIIntegrationService(api.APIIntegrationServiceOptions{Ctx: cmd.Context()})

	cluster, err := apiService.GetCluster(_describeCluster.clusterId)
	if err != nil {
		fmt.Println(cli.RenderError(fmt.Sprintf("There was an error describing cluster '%s': %s", _describeCluster.clusterId, err)) + "\n")
		return
	}

	s := &strings.Builder{}
	s.WriteString(cli.BoldStyle.Copy().Foreground(cli.Color).Padding(1, 0, 0, 0).Render(wrap.String("Describing your Cluster =>", width)))
	s.WriteString("\n")
	s.WriteString(cli.RenderAsJson(cluster))
	s.WriteString(cli.BoldStyle.Copy().Foreground(cli.Color).Padding(1, 0, 0, 0).Render(wrap.String("====================", width)))
	fmt.Println(s.String())
}
