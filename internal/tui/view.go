package tui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/dalin-stone/rescue-diver/internal/tui/components"
)

var (
	titleStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("205"))

	boxStyle = lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		Padding(1, 2).
		BorderForeground(lipgloss.Color("63"))
)

func (m Model) View() string {
	title := titleStyle.Render("Rescue Diver 🤿")
	body := fmt.Sprintf(
		"Typed: %s\n",
		m.input,
	)
	status := components.StatusBar(" Press Ctrl+c to quit")

	return lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		boxStyle.Render(body),
		status,
	)
}
