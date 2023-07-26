package common

import (
	"fmt"
	"strings"

	"github.com/muesli/reflow/wrap"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/onqlavelabs/onqlave.cli/internal/utils"
)

func CliRenderListResourceOutputNoRecord(width int) {
	s := &strings.Builder{}
	s.WriteString(utils.BoldStyle.Copy().Foreground(utils.Color).Padding(1, 0, 0, 0).Render(wrap.String("No record found", width)))
	s.WriteString("\n")
	fmt.Println(s.String())
}

func CliRenderListResourceOutput(width int, resource any, resourceName string) {
	s := &strings.Builder{}
	s.WriteString(utils.BoldStyle.Copy().Foreground(utils.Color).Padding(1, 0, 0, 0).Render(wrap.String(fmt.Sprintf("List %s =>", resourceName), width)))
	s.WriteString("\n")
	s.WriteString(utils.RenderAsJson(map[string]interface{}{resourceName: resource}))
	s.WriteString("\n")
	fmt.Println(s.String())
}

func CliRenderBaseResourceOutput(width int, resource any, resourceName string) {
	s := &strings.Builder{}
	s.WriteString(utils.BoldStyle.Copy().Foreground(utils.Color).Padding(1, 0, 0, 0).Render(wrap.String(fmt.Sprintf("%s Base Information =>", resourceName), width)))
	s.WriteString("\n")
	s.WriteString(utils.RenderAsJson(resource))
	s.WriteString("\n")
	fmt.Println(s.String())
}

func CliRenderSuccessActionResourceOutput(width int, resourceID, resourceName, action string) {
	s := &strings.Builder{}
	s.WriteString(utils.BoldStyle.Copy().Foreground(utils.Color).Padding(1, 0, 0, 0).Render(wrap.String(fmt.Sprintf("🎉 Done! %s %s successfully. \n", resourceName, action), width)))
	s.WriteString(utils.BoldStyle.Copy().Foreground(utils.Color).Padding(1, 0, 0, 0).Render(wrap.String(fmt.Sprintf("%s ID: %s", resourceName, resourceID), width)))
	s.WriteString(utils.BoldStyle.Copy().Foreground(utils.Color).Padding(1, 0, 0, 0).Render("For more information, read our documentation at https://docs.onqlave.com"))
	s.WriteString("\n")
	fmt.Println(s.String())
}

func CliRenderUIErrorOutput(ui *utils.SpinnerTUI, resourceName, actioned, resourceID string) {
	s := &strings.Builder{}
	if ui.Error() != nil {
		s.WriteString(utils.RenderError(fmt.Sprintf("There was an error whilst %s being %s: %s", resourceName, actioned, ui.Error())) + "\n")
	} else {
		s.WriteString(utils.BoldStyle.Copy().Foreground(utils.Green).Padding(1, 0, 0, 0).Render(fmt.Sprintf("🎉 Done! %s %s successfully. \n", resourceName, actioned)))
		s.WriteString(utils.BoldStyle.Copy().Foreground(utils.Color).Padding(1, 0, 0, 0).Render(fmt.Sprintf("%s ID: %s", resourceName, resourceID)))
		s.WriteString(utils.BoldStyle.Copy().Foreground(utils.Color).Padding(1, 0, 0, 0).Render("For more information, read our documentation at https://docs.onqlave.com"))
	}
	s.WriteString("\n")
	fmt.Println(s.String())
}

func CliRenderDescribeResourceOutput(width int, resource any, resourceName, resourceID string) {
	s := &strings.Builder{}
	s.WriteString(utils.BoldStyle.Copy().Foreground(utils.Color).Padding(1, 0, 0, 0).Render(wrap.String(fmt.Sprintf("%s '%s' Information =>", resourceName, resourceID), width)))
	s.WriteString("\n")
	s.WriteString(utils.RenderAsJson(map[string]interface{}{resourceName: resource}))
	s.WriteString("\n")
	fmt.Println(s.String())
}

func PersistentPreRun(cmd *cobra.Command, args []string) error {
	cmd.SilenceUsage = true

	if err := viper.BindPFlags(cmd.Flags()); err != nil {
		return ReplacePersistentPreRunE(cmd, err)
	}

	if !IsEnvConfigured() {
		return ReplacePersistentPreRunE(cmd, ErrUnsetEnv)
	}

	if !IsLoggedIn() {
		return ReplacePersistentPreRunE(cmd, ErrRequireLogIn)
	}

	cmd.SilenceUsage = false

	return nil
}

func ReplacePersistentPreRunE(cmd *cobra.Command, err error) error {
	cmd.SilenceErrors = true
	cmd.SilenceUsage = true
	fmt.Println(utils.RenderError(utils.BoldStyle.Render(fmt.Sprintf("%s", err))))
	return err
}
