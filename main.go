package main

import (
	tea "charm.land/bubbletea/v2"
	"elgca/cli/elgca/ui"
	"fmt"
	"log"
)

func main() {
	f, err := tea.LogToFile("debug.log", "")
	if err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
	}

	log.Printf("start")
	defer f.Close()
	p := tea.NewProgram(ui.NewMainMenu())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		return
	}
}
