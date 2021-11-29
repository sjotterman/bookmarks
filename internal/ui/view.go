package ui

import (
	"github.com/charmbracelet/lipgloss"
)

type Styles struct {
	ListStyle lipgloss.Style
}

func DefaultStyles() (s Styles) {
	s.ListStyle = lipgloss.NewStyle().Height(10).Padding(1, 0, 0, 2)
	return s
}

func (m Model) listView() string {
	return m.Styles.ListStyle.Render(m.list.View())
}

func (m Model) View() string {
	var (
		sections []string
	)
	sections = append(sections, m.listView())
	return lipgloss.JoinVertical(lipgloss.Left, sections...)
}
