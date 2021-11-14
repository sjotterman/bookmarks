package ui

func (m Model) View() string {
	return m.docStyle.Render(m.list.View())
}
