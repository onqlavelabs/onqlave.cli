package user

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/term"

	"github.com/onqlavelabs/onqlave.cli/cmd/common"
	"github.com/onqlavelabs/onqlave.cli/internal/pkg/cli/api/user/models"
)

func listCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "list",
		Short:   "list users",
		Long:    "This command is used to list cli existing users.",
		Example: "onqlave user list",
		Run:     runListCommand,
	}
}

func runListCommand(cmd *cobra.Command, args []string) {
	width, _, _ := term.GetSize(int(os.Stdout.Fd()))

	response, err := newUserApiService(cmd.Context()).GetPlatformOwnerAndClusterAdmin()
	if err != nil {
		common.RenderCLIOutputError("There was an error getting users: ", err)
		return
	}

	var userList models.UserList
	for _, u := range response.Users {
		userList.Users = append(userList.Users, u.User)
	}

	if len(userList.Users) == 0 {
		common.CliRenderListResourceOutputNoRecord(width)
		return
	}

	if viper.GetBool(common.FlagJson) {
		common.CliRenderListResourceOutput(width, userList.Users, common.ResourceUser)
		return
	}

	common.NewDataTable(userList.Users).Render()
}
