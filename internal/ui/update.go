package ui

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type updateBookmarkListMsg struct{ items []list.Item }

func (m *Model) handleXKeyPress(cmds *[]tea.Cmd) {
	m.list.Title = "Pressed a button!"
}

func (m *Model) handleUpdateBookmarksListMsg(msg updateBookmarkListMsg, cmds *[]tea.Cmd) {
	m.list.SetItems(msg.items)
}

// handleWindowSizeMsg is received whenever the window size changes.
func (m *Model) handleWindowSizeMsg(msg tea.WindowSizeMsg, cmds *[]tea.Cmd) {
	top, right, bottom, left := m.docStyle.GetMargin()
	m.list.SetSize(msg.Width-left-right, msg.Height-top-bottom)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.handleWindowSizeMsg(msg, &cmds)
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
		if msg.String() == "x" {
			m.handleXKeyPress(&cmds)
		}

	case updateBookmarkListMsg:
		m.handleUpdateBookmarksListMsg(msg, &cmds)

	}
	m.list, cmd = m.list.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}
