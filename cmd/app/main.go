package main

import(
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/dalin-stone/rescue-diver/internal/tui"
)

func main() {
	dock := tea.NewProgram(tui.NewModel())

	if err := dock.Start(); err != nil {
		log.Printf("Error Starting Program: %v\n", err)
		os.Exit(1)
	}
}
