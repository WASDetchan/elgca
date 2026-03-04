package main

import (
	tea "charm.land/bubbletea/v2"
	"elgca/cli/elgca/ui"
	"fmt"
)

func main() {
	p := tea.NewProgram(ui.NewMainMenu())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		return
	}
}
