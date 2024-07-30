// modified from <https://github.com/maaslalani/pom/blob/main/main.go>
package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/lipgloss"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

// init
const (
	titleColor   = "#2EF8BB"
	tickInterval = time.Second / 2
)

var (
	currentTime    = time.Now().Format("15:04PM")
	startTimeStyle = lipgloss.NewStyle().Foreground(lipgloss.Color(titleColor)).MarginRight(1).SetString(currentTime)
	helpStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("240")).MarginTop(2)
	baseTimerStyle = lipgloss.NewStyle().Padding(1, 2)
)

type tickMsg time.Time

// model
type Model struct {
	quitting  bool
	startTime time.Time
	duration  time.Duration
	progress  progress.Model
}

// bubbletea
func (m Model) Init() tea.Cmd {
	return tea.Tick(tickInterval, tickCmd)
}

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
			m.startTime = time.Now()

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
		cmds = append(cmds, tea.Tick(tickInterval, tickCmd))
	}

	if time.Since(m.startTime) > m.duration {
		m.startTime = time.Now()

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
	timeLeft := m.duration - elapsed

	percent := float64(elapsed) / float64(m.duration)
	s.WriteString(startTimeStyle.String())
	s.WriteString("- " + timeLeft.Round(time.Second).String())
	s.WriteString("\n\n")
	s.WriteString(m.progress.ViewAs(percent))
	s.WriteString(helpStyle.Render("Press 'q' to quit"))

	return baseTimerStyle.Render(s.String())
}

func NewModel() Model {
	progressBar := progress.New()
	progressBar.FullColor = titleColor
	progressBar.SetSpringOptions(1, 1)

	return Model{
		progress: progressBar,
	}
}

// main
var TimerCmd = &cobra.Command{
	Use:   "timer",
	Short: "Create a timer",
	Run: func(cmd *cobra.Command, args []string) {
		// check if input exists
		if len(args) == 0 {
			fmt.Println("Please provide a duration")
			os.Exit(1)
		}
		// parse input
		duration, err := time.ParseDuration(args[0])
		if err != nil {
			log.Fatal(err)
		}

		// timer
		m := NewModel()
		m.duration = duration

		_, err = tea.NewProgram(&m).Run()
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(TimerCmd)
}
