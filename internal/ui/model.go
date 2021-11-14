package ui

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

type item struct {
	title, desc string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

type Model struct {
	docStyle lipgloss.Style
	list         list.Model
	saveFilePath string
}

func NewModel(saveFilePath string) Model {
	items := []list.Item{}
	return Model{
		docStyle:     lipgloss.NewStyle().Margin(1, 0, 0, 0),
		list:         list.NewModel(items, list.NewDefaultDelegate(), 0, 0),
		saveFilePath: saveFilePath,
	}
}
