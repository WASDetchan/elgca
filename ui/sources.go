package ui

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

var (
	SourcesStyle = lipgloss.NewStyle().
			Width(32).
			Height(5).
			Align(lipgloss.Left, lipgloss.Top).
			Padding(1).
			Margin(0).
			Border(lipgloss.NormalBorder(), false, true, false, false).
			BorderForeground(lipgloss.Color("69"))
	TitleStyle        = lipgloss.NewStyle().MarginLeft(2)
	ItemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	SelectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
)

type sourceItem struct {
	name       string
	sourceType string
	component  Component
}

type Sources struct {
	selectedSource int
	sources        []sourceItem
}

func NewSources() Sources {
	return Sources{
		0,
		[]sourceItem{},
	}
}

func (m Sources) Init() tea.Cmd {
	return nil
}

var creation SourceCreation = NewSourceCreation()

func (m Sources) Update(msg tea.Msg) (Component, tea.Cmd) {
	propagate := true
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "J":
			propagate = false
			if m.selectedSource < len(m.sources) {
				m.selectedSource++
			}
		case "K":
			propagate = false
			if m.selectedSource > 0 {
				m.selectedSource--
			}
		}
	}

	innerCmd := tea.Cmd(nil)

	if propagate && m.selectedSource < len(m.sources) {
		m.sources[m.selectedSource].component, innerCmd =
			m.sources[m.selectedSource].component.Update(msg)
	}
	if m.selectedSource == len(m.sources) {
		if propagate {
			creation, innerCmd = creation.Update(msg)
		}

		newSource, finished := creation.Finish()

		if finished {
			creation = NewSourceCreation()
			m.sources = append(m.sources, *newSource)
		}
	}

	return m, innerCmd
}

func (m Sources) Render() string {
	var cont strings.Builder

	for idx, s := range m.sources {
		selectedMark := " "
		if m.selectedSource == idx {
			selectedMark = ">"
		}
		fmt.Fprintf(&cont, "%s %s\n", selectedMark, s.name)
	}

	sourceContent := ""
	selectedMark := " "
	if m.selectedSource == len(m.sources) {
		selectedMark = ">"
		sourceContent = creation.Render()
	} else {
		sourceContent = m.sources[m.selectedSource].component.Render()
	}
	fmt.Fprintf(&cont, "%s %s\n", selectedMark, "new source")

	selfContent := SourcesStyle.Render(cont.String())

	content := lipgloss.JoinHorizontal(lipgloss.Top,
		selfContent,
		sourceContent,
	)

	return content
}
