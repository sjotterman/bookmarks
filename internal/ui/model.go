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
	list     list.Model
}

func NewModel() Model {
	items := []list.Item{
		item{title: "foo.md", desc: "First file"},
		item{title: "bar.md", desc: "Second file"},
		item{title: "foobar.md", desc: "Third file"},
		item{title: "baz.md", desc: "Fourth file"},
	}
	return Model{
		docStyle: lipgloss.NewStyle().Margin(1, 0, 0, 0),
		list:     list.NewModel(items, list.NewDefaultDelegate(), 0, 0),
	}
}
