package ui

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type section struct {
	name      string
	component Component
}

type MainMenu struct {
	selectedSection int
	sections        []section
}

var (
	SidebarStyle = lipgloss.NewStyle().
		Width(16).
		Height(5).
		Padding(1).
		Align(lipgloss.Left, lipgloss.Center).
		Border(lipgloss.NormalBorder(), false, true, false, false).
		BorderForeground(lipgloss.Color("69"))
)

func updateStyles(msg tea.Msg) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		SidebarStyle = SidebarStyle.Height(msg.Height)
		HomeStyle = HomeStyle.Height(msg.Height).Width(msg.Width - SidebarStyle.GetWidth())
		SourcesStyle = SourcesStyle.Height(msg.Height)
		EmptySourceStyle = EmptySourceStyle.
			Height(msg.Height).
			Width(msg.Width -
				SidebarStyle.GetWidth() -
				SourcesStyle.GetWidth(),
			)
		SourceCreationStyle = SourceCreationStyle.
			Height(msg.Height).
			Width(msg.Width -
				SidebarStyle.GetWidth() -
				SourcesStyle.GetWidth(),
			)
	}
}

func NewMainMenu() MainMenu {
	return MainMenu{0, []section{
		{"home", NewHome()},
		{"playlists", NewHome()},
		{"sources", NewSources()},
		{"settings", NewHome()},
	}}
}

func startLog() tea.Msg {
	return nil
}

func (m MainMenu) Init() tea.Cmd {
	return nil
}

func (m MainMenu) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	propagate := true
	updateStyles(msg)
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "j":
			propagate = false
			if m.selectedSection+1 < len(m.sections) {
				m.selectedSection++
			}
		case "k":
			propagate = false
			if m.selectedSection > 0 {
				m.selectedSection--
			}
		}
	}

	innerCmd := tea.Cmd(nil)

	if propagate {
		m.sections[m.selectedSection].component, innerCmd =
			m.sections[m.selectedSection].component.Update(msg)
	}

	return m, innerCmd
}

func (m MainMenu) View() tea.View {
	var cont strings.Builder

	for idx, s := range m.sections {
		selectedMark := " "
		if m.selectedSection == idx {
			selectedMark = ">"
		}
		fmt.Fprintf(&cont, "%s %s\n", selectedMark, s.name)
	}

	selfContent := SidebarStyle.Render(cont.String())

	content := lipgloss.JoinHorizontal(lipgloss.Top,
		selfContent,
		m.sections[m.selectedSection].component.Render(),
	)

	view := tea.NewView(content)

	view.AltScreen = true
	return view
}
