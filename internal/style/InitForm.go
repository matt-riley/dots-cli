package style

import (
	"github.com/charmbracelet/lipgloss"
)

// InitFormStyles holds the styles for the init form
type InitFormStyles struct {
	BorderColour lipgloss.Color
	InputField   lipgloss.Style
	Question     lipgloss.Style
	DoneMsg      lipgloss.Style
	DoneExit     lipgloss.Style
}

// DefaultInitFormStyles returns the default styles for the init form
func DefaultInitFormStyles() *InitFormStyles {
	s := new(InitFormStyles)
	s.BorderColour = Mauve
	s.InputField = lipgloss.NewStyle().
		BorderForeground(s.BorderColour).
		BorderStyle(lipgloss.RoundedBorder()).
		Padding(0, 1).
		Width(80)
	s.Question = lipgloss.NewStyle().Bold(true).Foreground(Green)
	s.DoneMsg = lipgloss.NewStyle().Foreground(Red).Bold(true)
	s.DoneExit = lipgloss.NewStyle().Foreground(Text)
	return s
}
