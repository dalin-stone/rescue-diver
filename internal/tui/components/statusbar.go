package components

import "github.com/charmbracelet/lipgloss"

var statusBarStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("0")).
	Background(lipgloss.Color("63")).
	Padding(0, 1)

func StatusBar(text string) string {
	return statusBarStyle.Render(text)
}

