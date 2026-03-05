package ui

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

var (
	HomeStyle = lipgloss.NewStyle().
		Width(32).
		Height(5).
		Align(lipgloss.Center, lipgloss.Top).
		Padding(0).
		Margin(0).
		BorderStyle(lipgloss.HiddenBorder())
)

type Home struct {
}

func NewHome() Home {
	return Home{}
}

func (m Home) Init() tea.Cmd {
	return nil
}

func (m Home) Update(msg tea.Msg) (Component, tea.Cmd) {
	return m, nil
}

func (m Home) View() tea.View {
	return tea.NewView(m.Render())
}

func (m Home) Render() string {
	return HomeStyle.Render("hi!")
}
