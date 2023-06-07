package cli

import (
	"encoding/json"
	"strings"

	"github.com/TylerBrock/colorjson"
	"github.com/charmbracelet/lipgloss"
	"github.com/gosimple/slug"
)

var (
	Color   = lipgloss.AdaptiveColor{Light: "#111222", Dark: "#FAFAFA"}
	Primary = lipgloss.Color("#4636f5")
	Green   = lipgloss.Color("#9dcc3a")
	Red     = lipgloss.Color("#ff0000")
	White   = lipgloss.Color("#ffffff")
	Black   = lipgloss.Color("#000000")
	Orange  = lipgloss.Color("#D3A347")
	Feint   = lipgloss.AdaptiveColor{Light: "#333333", Dark: "#888888"}
	Iris    = lipgloss.Color("#5D5FEF")
	Fuschia = lipgloss.Color("#EF5DA8")

	TextStyle  = lipgloss.NewStyle().Foreground(Color)
	FeintStyle = TextStyle.Copy().Foreground(Feint)
	BoldStyle  = TextStyle.Copy().Bold(true)
)

// RenderError returns a formatted error string.
func RenderError(msg string) string {
	// Error applies styles to an error message
	err := lipgloss.NewStyle().Background(Red).Foreground(White).Bold(true).Padding(0, 1).Render("Error")
	content := lipgloss.NewStyle().Bold(true).Padding(0, 1).Render(msg)
	return err + content
}

// RenderWarning returns a formatted warning string.
func RenderWarning(msg string) string {
	// Error applies styles to an error message
	err := lipgloss.NewStyle().Foreground(Orange).Bold(true).Render("Warning: ")
	content := lipgloss.NewStyle().Bold(true).Foreground(Orange).Padding(0, 1).Render(msg)
	return err + content
}

func Slug(name string) string {
	return strings.ToLower(slug.Make(name))
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
