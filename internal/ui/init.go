package ui

import (
	tea "github.com/charmbracelet/bubbletea"
)

// Initialize UI and set up initial data
// Right now this doesn't do anything, but this is
// where the initial bookmarks should be read in
func (m Model) Init() tea.Cmd {
	return m.loadBookmarksCmd()
}
