package ui

import (
	"log"

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

var (
	SourceCreationStyle = lipgloss.NewStyle().
		Width(32).
		Height(5).
		Align(lipgloss.Left, lipgloss.Top).
		Padding(0).
		Margin(0).
		BorderStyle(lipgloss.HiddenBorder())
)

type SourceCreation struct {
	name      string
	nameInput textinput.Model
	provider  string
}

func NewSourceCreation() SourceCreation {
	nameInput := textinput.New()
	nameInput.Placeholder = "Source Name"
	nameInput.SetVirtualCursor(false)
	nameInput.Focus()
	nameInput.CharLimit = 64
	nameInput.SetWidth(64)

	return SourceCreation{
		"",
		nameInput,
		"",
	}
}

func (m SourceCreation) Finish() (*sourceItem, bool) {
	if m.name == "" {
		return nil, false
	}
	source := NewEmptySource()
	source.name = m.name
	return &source, true
}

func (m SourceCreation) Update(msg tea.Msg) (SourceCreation, tea.Cmd) {
	var cmd tea.Cmd
	if m.name == "" {
		keymsg, ok := msg.(tea.KeyPressMsg)
		if ok {
			switch keymsg.String() {
			case "enter":
				m.name = m.nameInput.Value()
				log.Printf("Set name: %s", m.name)
			case "l":
				m.nameInput.Focus()
			}
		}
	}
	if m.name == "" {
		m.nameInput, cmd = m.nameInput.Update(msg)
	}
	return m, cmd
}

func (m SourceCreation) Render() string {
	name := m.name
	if name == "" {
		name = m.nameInput.View()
	}
	return lipgloss.JoinVertical(lipgloss.Center, "Name: "+name)
}
