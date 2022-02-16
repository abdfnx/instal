package tui

import (
	"os"
	"fmt"
	"runtime"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/abdfnx/instal/core/options"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/abdfnx/instal/core/installer"
	"github.com/charmbracelet/bubbles/textinput"
)

var (
	focusedStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	blurredStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	cursorStyle         = focusedStyle.Copy()
	noStyle             = lipgloss.NewStyle()

	focusedButton = focusedStyle.Copy().Render("[ OK ]")
	blurredButton = fmt.Sprintf("[ %s ]", blurredStyle.Render("OK"))
)

type model struct {
	focusIndex int
	inputs     []textinput.Model
}

func initialModel() model {
	m := model{
		inputs: make([]textinput.Model, 3),
	}

	p := "bash"

	if runtime.GOOS == "windows" {
		p = "powershell.exe"
	}

	var t textinput.Model
	for i := range m.inputs {
		t = textinput.New()
		t.CursorStyle = cursorStyle
		t.CharLimit = 32

		switch i {
			case 0:
				t.Placeholder = "URL"
				t.Focus()
				t.PromptStyle = focusedStyle
				t.TextStyle = focusedStyle

			case 1:
				t.Placeholder = "Shell to use (" + p + ")"
				t.CharLimit = 64

			case 2:
				t.Placeholder = "Is hidden (y/n)"
				t.CharLimit = 5
		}

		m.inputs[i] = t
	}

	return m
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
				case "ctrl+c", "esc":
					return m, tea.Quit

				// Set focus to next input
				case "tab", "shift+tab", "enter", "up", "down":
					s := msg.String()

					if s == "enter" && m.focusIndex == len(m.inputs) {
						opts := options.InstalOptions{
							Shell: "",
							IsHidden: false,
							URL: "",
						}

						installer.RunInstal(&opts, false, m.inputs[0].Value(), m.inputs[1].Value(), m.inputs[2].Value())

						return m, tea.Quit
					}

					// Cycle indexes
					if s == "up" || s == "shift+tab" {
						m.focusIndex--
					} else {
						m.focusIndex++
					}

					if m.focusIndex > len(m.inputs) {
						m.focusIndex = 0
					} else if m.focusIndex < 0 {
						m.focusIndex = len(m.inputs)
					}

					cmds := make([]tea.Cmd, len(m.inputs))

					for i := 0; i <= len(m.inputs)-1; i++ {
						if i == m.focusIndex {
							// Set focused state
							cmds[i] = m.inputs[i].Focus()
							m.inputs[i].PromptStyle = focusedStyle
							m.inputs[i].TextStyle = focusedStyle
							continue
						}

						// Remove focused state
						m.inputs[i].Blur()
						m.inputs[i].PromptStyle = noStyle
						m.inputs[i].TextStyle = noStyle
					}

					return m, tea.Batch(cmds...)
			}
	}

	// Handle character input and blinking
	cmd := m.updateInputs(msg)

	return m, cmd
}

func (m *model) updateInputs(msg tea.Msg) tea.Cmd {
	var cmds = make([]tea.Cmd, len(m.inputs))

	// Only text inputs with Focus() set will respond, so it's safe to simply
	// update all of them here without any further logic.
	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}

func (m model) View() string {
	var b strings.Builder

	logo := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#fff")).
		Background(lipgloss.Color("#6957B0")).
		Padding(0, 1).
		SetString("Instal")

	paddingLogo := lipgloss.NewStyle().
		Padding(0, 2).
		SetString(logo.String())

	b.WriteString("\n" + paddingLogo.String() + "\n\n")

	for i := range m.inputs {
		b.WriteString(lipgloss.NewStyle().Padding(0, 2).SetString(m.inputs[i].View()).String() + "\n")
	}

	button := &blurredButton

	if m.focusIndex == len(m.inputs) {
		button = &focusedButton
	}

	fmt.Fprintf(&b, "\n%s\n\n", lipgloss.NewStyle().Padding(0, 2).SetString(*button).String())

	return b.String()
}

func Instal() {
	if err := tea.NewProgram(initialModel()).Start(); err != nil {
		fmt.Printf("could not start program: %s\n", err)
		os.Exit(1)
	}
}
