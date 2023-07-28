package utils

import (
	"encoding/json"

	"github.com/TylerBrock/colorjson"
	"github.com/charmbracelet/lipgloss"
)

var (
	Color   = lipgloss.AdaptiveColor{Light: "#111222", Dark: "#FAFAFA"}
	Primary = lipgloss.Color("#4636f5")
	Green   = lipgloss.Color("#9dcc3a")
	Red     = lipgloss.Color("#ff0000")
	White   = lipgloss.Color("#ffffff")

	TextStyle = lipgloss.NewStyle().Foreground(Color)
	BoldStyle = TextStyle.Copy().Bold(true)
)

// RenderError returns a formatted error string.
func RenderError(msg string) string {
	// Error applies styles to an error message
	err := lipgloss.NewStyle().Background(Red).Foreground(White).Bold(true).Padding(0, 1).Render("Error")
	content := lipgloss.NewStyle().Bold(true).Padding(0, 1).Render(msg)
	return err + content
}
func RenderAsJson(object interface{}) string {
	var objMap map[string]interface{}
	marshalledObj, _ := json.Marshal(object)
	_ = json.Unmarshal(marshalledObj, &objMap)

	f := colorjson.NewFormatter()
	f.Indent = 4
	bytes, _ := f.Marshal(objMap)
	return string(bytes)
}
