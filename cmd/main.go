package main

import (
	"fmt"
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	choices  []string
	cursor   int
	selected map[int]struct{}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch message := msg.(type) {
	case tea.KeyMsg:
		switch message.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor = m.cursor - 1
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor = m.cursor + 1
			}
		case "enter", " ":
			if _, ok := m.selected[m.cursor]; ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}

	}
	return m, nil
}

func (m model) View() string {
	s := "A simple choice selector.\n\n"
	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		checked := ""
		if _, ok := m.selected[i]; ok {
			checked = "x"
		}
		s = s + fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}
	return s
}

func NewModel() model {
	return model{
		choices:  []string{"foo", "bar", "baz"},
		cursor:   0,
		selected: make(map[int]struct{}),
	}
}

func main() {
	p := tea.NewProgram(NewModel())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
