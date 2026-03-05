package ui

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

var (
	EmptySourceStyle = lipgloss.NewStyle().
		Width(32).
		Height(5).
		Align(lipgloss.Left, lipgloss.Top).
		Padding(0).
		Margin(0).
		BorderStyle(lipgloss.HiddenBorder())
)

type EmptySource struct {
}

func NewEmptySource() sourceItem {
	return sourceItem{
		"Empty Source",
		"empty",
		EmptySource{},
	}
}

func (m EmptySource) Init() tea.Cmd {
	return nil
}

func (m EmptySource) Update(msg tea.Msg) (Component, tea.Cmd) {
	return m, nil
}

func (m EmptySource) View() tea.View {
	return tea.NewView(m.Render())
}

func (m EmptySource) Render() string {
	return EmptySourceStyle.Render("nothin her yet")
}
