package tui

import (
	"net/url"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/matt-riley/dots-cli/internal/style"
)

// InitPromptModel is a model for the init prompt
type InitPromptModel struct {
	exists    bool
	index     int
	width     int
	height    int
	questions []InitQuestion
	styles    *style.InitFormStyles
	done      bool
	validURL  bool
}

// InitQuestion is a question for the init prompt
type InitQuestion struct {
	question string
	answer   string
	input    Input
	url      bool
}

// Next increments the index of the init prompt
func (m *InitPromptModel) Next() {
	if m.index < len(m.questions)-1 {
		m.index++
	} else {
		m.index = 0
	}
}
func newQuestion(q string) InitQuestion {
	return InitQuestion{question: q}
}

func newShortQuestion(q string, url bool) InitQuestion {
	question := newQuestion(q)
	model := NewShortAnswerField()
	question.input = model
	question.url = url
	return question
}

// InitialInitModel returns an initial model for the init prompt
func InitialInitModel(exists bool) InitPromptModel {
	questions := []InitQuestion{
		newShortQuestion("The URL of the dotfiles repo", true),
		newShortQuestion("The branch to use", false),
		newShortQuestion("Your Github PAT", false),
	}
	styles := style.DefaultInitFormStyles()
	mod := InitPromptModel{}
	mod.exists = exists
	mod.questions = questions
	mod.styles = styles
	mod.validURL = true
	return mod
}

func validateURL(input string) bool {
	_, err := url.ParseRequestURI(input)
	return err == nil
}

// Init returns a command for the init prompt
func (m InitPromptModel) Init() tea.Cmd {
	return m.questions[m.index].input.Blink
}

// Update updates the init prompt model
func (m InitPromptModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	currentQuestion := &m.questions[m.index]
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case tea.KeyMsg:
		switch msg.String() {
		case "esc", "ctrl+c":
			return m, tea.Quit
		case "enter":
			if m.done {
				return m, tea.Quit
			}
			if currentQuestion.url {
				valid := validateURL(currentQuestion.input.Value())
				m.validURL = valid
			} else {
				m.validURL = true
			}
			if m.index == len(m.questions)-1 {
				m.done = true
			}

			if m.validURL {
				currentQuestion.answer = currentQuestion.input.Value()
				m.Next()
			}
			return m, currentQuestion.input.Blur
		}
	}
	currentQuestion.input, cmd = currentQuestion.input.Update(msg)
	return m, cmd
}

// View returns a string for the init prompt
func (m InitPromptModel) View() string {
	if m.width == 0 {
		return "loading"
	}
	if m.done {
		return lipgloss.Place(
			m.width,
			m.height,
			lipgloss.Center,
			lipgloss.Center,
			lipgloss.JoinVertical(
				lipgloss.Center,
				m.styles.DoneMsg.Render("Done!"),
				m.styles.DoneExit.Render("Press enter to exit"),
			),
		)
	}
	current := m.questions[m.index]

	if !m.validURL {
		return lipgloss.Place(
			m.width,
			m.height,
			lipgloss.Center,
			lipgloss.Center,
			lipgloss.JoinVertical(
				lipgloss.Left,
				m.styles.Question.Render(
					m.questions[m.index].question,
				)+" "+lipgloss.NewStyle().
					Bold(true).
					Foreground(style.Red).
					Render("Please amend input to be a valid URL"),
				m.styles.InputField.BorderForeground(style.Red).
					Render(current.input.View()),
			),
		)
	}

	return lipgloss.Place(
		m.width, m.height,
		lipgloss.Center, lipgloss.Center,
		lipgloss.JoinVertical(
			lipgloss.Left,
			m.styles.Question.Render(m.questions[m.index].question),
			m.styles.InputField.BorderForeground(style.Green).
				Render(current.input.View()),
		),
	)
}
