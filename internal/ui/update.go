package ui

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (m *Model) handleXKeyPress(cmds *[]tea.Cmd) {
	m.text = "Updated!"
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
		if msg.String() == "x" {
			m.handleXKeyPress(&cmds)
		}

	}
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}
