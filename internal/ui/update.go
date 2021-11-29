package ui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type updateBookmarkListMsg struct{ items []item }
type saveFileNotExistsMsg struct{ filename string }
type emptyFileCreatedMsg struct{ filename string }
type markItemReadMsg struct{ filename string }

func (m *Model) handleToggleRead(msg tea.Msg, cmds *[]tea.Cmd) {
	index := m.list.Index()
	newItem := m.items[index]
	newItem.isRead = !newItem.isRead
	m.items[index] = newItem
	*cmds = append(*cmds, m.list.SetItem(index, newItem))
}

func (m *Model) handleUpdateBookmarksListMsg(msg updateBookmarkListMsg, cmds *[]tea.Cmd) {
	cmd := m.list.NewStatusMessage("Loaded bookmarks!")
	*cmds = append(*cmds, cmd)
	m.items = msg.items
	var listItems []list.Item
	// This is kind of hacky, but it's the only way I can find to
	// access and modify the items in place
	for _, item := range msg.items {
		listItems = append(listItems, item)
	}
	*cmds = append(*cmds, m.list.SetItems(listItems))
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
		switch {
		case key.Matches(msg, m.KeyMap.ToggleRead):
			m.handleToggleRead(msg, &cmds)
		}
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
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
