package ui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type updateBookmarkListMsg struct{ items []list.Item }
type saveFileNotExistsMsg struct{ filename string }
type emptyFileCreatedMsg struct{ filename string }

func (m *Model) handleXKeyPress(cmds *[]tea.Cmd) {
	m.list.Title = "Pressed a button!"
}

func (m *Model) handleUpdateBookmarksListMsg(msg updateBookmarkListMsg, cmds *[]tea.Cmd) {
	cmd := m.list.NewStatusMessage("Loaded bookmarks!")
	*cmds = append(*cmds, cmd)
	m.list.SetItems(msg.items)
}

func (m *Model) handleCreateEmptyFileMsg(msg saveFileNotExistsMsg, cmds *[]tea.Cmd) {
	*cmds = append(*cmds, m.createEmptySaveFileCmd(msg.filename))
}

func (m *Model) handleEmptyFileCreatedMsg(msg emptyFileCreatedMsg, cmds *[]tea.Cmd) {
	statusMsg := fmt.Sprintf("Created %s", msg.filename)
	cmd := m.list.NewStatusMessage(statusMsg)
	*cmds = append(*cmds, cmd)
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
	case saveFileNotExistsMsg:
		m.handleCreateEmptyFileMsg(msg, &cmds)
	case emptyFileCreatedMsg:
		m.handleEmptyFileCreatedMsg(msg, &cmds)

	}
	m.list, cmd = m.list.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}
