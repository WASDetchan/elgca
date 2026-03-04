package main

import (
	"fmt"
	"strings"

	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type item struct {
	value int
}

func (i item) Title() string       { return fmt.Sprint(i.value) }
func (i item) Description() string { return "" }
func (i item) FilterValue() string { return fmt.Sprint(i.value) }

type model struct {
	list   list.Model
	cursor int
}

func newModel() model {
	return model{list.New([]list.Item{item{0}}, list.NewDefaultDelegate(), 0, 0), 0}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrt+c", "q":
			return m, tea.Quit
		case "down", "j":
			if m.cursor+1 < len(m.numbers) {
				m.cursor++
			}
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "enter":
			m.numbers = append(m.numbers, m.numbers[m.cursor])
			m.cursor = len(m.numbers) - 1
		case "+":
			m.numbers[m.cursor]++
		}
	}
	return m, nil
}

func (m model) View() tea.View {
	var s strings.Builder
	s.WriteString("EXPAND!\n\n")

	for i, number := range m.numbers {

		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		// Render the row
		fmt.Fprintf(&s, "%s %d\n", cursor, number)
	}

	s.WriteString("\nPress q to quit.\n")

	v := tea.NewView(docStyle.Render(m.list.View()))
	v.AltScreen = true

	return tea.NewView(s.String())
}

func main() {
	p := tea.NewProgram(newModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		return
	}
}
