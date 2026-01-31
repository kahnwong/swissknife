package timer

import (
	"testing"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

func TestNewModel(t *testing.T) {
	m := NewModel()

	if m.progress.FullColor != titleColor {
		t.Errorf("NewModel() progress color = %v, want %v", m.progress.FullColor, titleColor)
	}

	if m.duration != 0 {
		t.Errorf("NewModel() duration = %v, want 0", m.duration)
	}

	if !m.startTime.IsZero() {
		t.Errorf("NewModel() startTime should be zero")
	}
}

func TestTimerValidation(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		wantErr bool
	}{
		{
			name:    "no args",
			args:    []string{},
			wantErr: true,
		},
		{
			name:    "invalid duration",
			args:    []string{"invalid"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Timer(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("Timer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestModelInit(t *testing.T) {
	m := NewModel()
	cmd := m.Init()

	if cmd == nil {
		t.Error("Init() should return a command")
	}
}

func TestModelUpdate(t *testing.T) {
	m := NewModel()
	m.duration = 5 * time.Second
	m.startTime = time.Now()

	t.Run("tick message", func(t *testing.T) {
		newModel, cmd := m.Update(tickMsg(time.Now()))
		if cmd == nil {
			t.Error("Update with tickMsg should return a command")
		}
		if newModel.(Model).quitting {
			t.Error("Model should not be quitting after tick")
		}
	})

	t.Run("quit key", func(t *testing.T) {
		newModel, _ := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'q'}})
		if !newModel.(Model).quitting {
			t.Error("Model should be quitting after 'q' key")
		}
	})

	t.Run("ctrl+c", func(t *testing.T) {
		newModel, _ := m.Update(tea.KeyMsg{Type: tea.KeyCtrlC})
		if !newModel.(Model).quitting {
			t.Error("Model should be quitting after ctrl+c")
		}
	})
}

func TestModelView(t *testing.T) {
	m := NewModel()
	m.duration = 5 * time.Second
	m.startTime = time.Now()

	t.Run("normal view", func(t *testing.T) {
		view := m.View()
		if view == "" {
			t.Error("View() should return non-empty string when not quitting")
		}
	})

	t.Run("quitting view", func(t *testing.T) {
		m.quitting = true
		view := m.View()
		if view != "" {
			t.Error("View() should return empty string when quitting")
		}
	})
}

func TestTickCmd(t *testing.T) {
	now := time.Now()
	msg := tickCmd(now)

	if _, ok := msg.(tickMsg); !ok {
		t.Error("tickCmd() should return tickMsg type")
	}
}
