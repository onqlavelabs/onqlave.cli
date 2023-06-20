package cli

import (
	"context"
	"errors"
	"fmt"
	"github.com/onqlavelabs/onqlave.cli/internal/cli/api"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type SpinnerTUI struct {
	duration  time.Duration
	options   SpinnerOptions
	start     bool
	waiting   spinner.Model
	startTime time.Time
	result    api.ConcurrencyOperationResult
}

type SpinnerOptions struct {
	Valid    int
	Consumer *api.Consumer
}

func NewSpnnerTUI(ctx context.Context, opts SpinnerOptions) (*SpinnerTUI, error) {
	r := &SpinnerTUI{
		options: opts,
		waiting: spinner.New(),
		result: api.ConcurrencyOperationResult{
			Result: nil,
			Done:   false,
			Error:  nil,
		},
	}
	r.waiting.Spinner = spinner.Dot
	r.waiting.Style = lipgloss.NewStyle().Foreground(Primary)

	return r, nil
}

func (ui *SpinnerTUI) Error() error {
	return ui.result.Error
}

func (ui *SpinnerTUI) Init() tea.Cmd {
	go ui.runOperation()

	return spinner.Tick
}

func (ui *SpinnerTUI) runOperation() {
	ui.startTime = time.Now().UTC()
	ui.start = true
	ui.options.Consumer.Consume(func(result api.ConcurrencyOperationResult) {
		ui.result = result
	})
	ui.duration = time.Since(ui.startTime)
	ui.result.Done = true
}

func (ui *SpinnerTUI) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	// Enable quitting early.
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyTab, tea.KeyShiftTab, tea.KeyEnter, tea.KeyUp, tea.KeyDown:
		case tea.KeyCtrlC, tea.KeyCtrlBackslash:
			ui.result.Error = errors.New("Program quit unexpectedly")
			return ui, tea.Quit
		}
		if msg.String() == "q" {
			ui.result.Error = errors.New("Program quit unexpectedly")
			return ui, tea.Quit
		}
	}

	ui.waiting, cmd = ui.waiting.Update(msg)
	cmds = append(cmds, cmd)
	if ui.result.Done || ui.result.Error != nil || ui.duration != 0 {
		cmds = append(cmds, tea.Quit)
	}

	return ui, tea.Batch(cmds...)
}

func (ui *SpinnerTUI) View() string {
	s := &strings.Builder{}

	s.WriteString(ui.RenderState())

	return s.String()
}

func (ui *SpinnerTUI) RenderState() string {
	if !ui.result.Done && ui.result.Error == nil {
		return fmt.Sprintf("%s %s... Please be patient 🥳\n\n", ui.waiting.View(), ui.result.Result)
	} else {
		return string("")
	}
}
