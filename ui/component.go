package ui

import tea "charm.land/bubbletea/v2"

type Component interface {
	Init() tea.Cmd
	Update(tea.Msg) (Component, tea.Cmd)
	Render() string
}
