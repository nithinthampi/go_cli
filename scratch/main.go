package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	lipgloss "github.com/charmbracelet/lipgloss"
)

type model struct {
	message string
}

var helpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("240")).Render

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		{
			switch msg.String() {
			case "q", "ctrl+c":
				return m, tea.Quit
			}
		}
	}
	return m, nil
}

func (m model) View() string {
	return helpStyle(m.message)
}

func main() {
	if _, err := tea.NewProgram(model{message: "foo bar"}).Run(); err != nil {
		log.Fatal(err)
	}

}
