package misc

import (
	"log"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/lipgloss"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

const (
	focusColor = "#2EF8BB"
	breakColor = "#FF5F87"
)

var (
	focusTitleStyle = lipgloss.NewStyle().Foreground(lipgloss.Color(focusColor)).MarginRight(1).SetString("Focus Mode") // [TODO] change me
	helpStyle       = lipgloss.NewStyle().Foreground(lipgloss.Color("240")).MarginTop(2)
)

var baseTimerStyle = lipgloss.NewStyle().Padding(1, 2)

type mode int

const (
	Initial mode = iota
	Focusing
	Paused
)

type Model struct {
	quitting bool

	startTime time.Time

	mode mode

	focusTime time.Duration

	progress progress.Model
}

func (m Model) Init() tea.Cmd {
	return tea.Tick(tickInterval, tickCmd)
}

const tickInterval = time.Second / 2

type tickMsg time.Time

func tickCmd(t time.Time) tea.Msg {
	return tickMsg(t)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tickMsg:
		cmds = append(cmds, tea.Tick(tickInterval, tickCmd))
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			m.mode = Paused
			m.startTime = time.Now()
			m.progress.FullColor = breakColor

			m.quitting = true
			return m, tea.Quit
		case "ctrl+c":
			m.quitting = true
			return m, tea.Quit
		default:
			m.startTime = time.Now()
		}
	}

	// Update timer
	if m.startTime.IsZero() {
		m.startTime = time.Now()
		m.mode = Focusing
		cmds = append(cmds, tea.Tick(tickInterval, tickCmd))
	}

	if time.Since(m.startTime) > m.focusTime {
		m.mode = Paused
		m.startTime = time.Now()
		m.progress.FullColor = breakColor

		m.quitting = true
		return m, tea.Quit
	}

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	if m.quitting {
		return ""
	}

	var s strings.Builder

	elapsed := time.Since(m.startTime)

	var percent float64
	switch m.mode {
	case Focusing:
		percent = float64(elapsed) / float64(m.focusTime)
		s.WriteString(focusTitleStyle.String())
		s.WriteString(elapsed.Round(time.Second).String())
		s.WriteString("\n\n")
		s.WriteString(m.progress.ViewAs(percent))
		s.WriteString(helpStyle.Render("Press 'q' to quit"))
	}

	return baseTimerStyle.Render(s.String())
}

func NewModel() Model {
	progressBar := progress.New()
	progressBar.FullColor = focusColor
	progressBar.SetSpringOptions(1, 1)

	return Model{
		progress: progressBar,
	}
}

var TimerCmd = &cobra.Command{
	Use:   "timer",
	Short: "Create a timer",
	Run: func(cmd *cobra.Command, args []string) {
		m := NewModel()

		m.focusTime = time.Duration(5 * float64(time.Second))

		_, err := tea.NewProgram(&m).Run()
		if err != nil {
			log.Fatal(err)
		}
	},
}
