package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	input string
}

func NewModel() Model {
	return Model{
		input: "",
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}
