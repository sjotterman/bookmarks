package ui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

type item struct {
	title, desc string
	isRead      bool
}

func (i item) Title() string {
	if !i.isRead {
		return fmt.Sprintf("☐ %s", i.title)
	}
	return fmt.Sprintf("☑ %s", i.title)
}

func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

type Model struct {
	docStyle lipgloss.Style
	Styles   Styles
	list         list.Model
	items        []item
	saveFilePath string
	KeyMap       *KeyMap
}

func NewModel(saveFilePath string) Model {
	var keyMap = newKeyMap()
	items := []list.Item{}
	m := Model{
		docStyle:     lipgloss.NewStyle().Margin(1, 0, 0, 0),
		list:         list.NewModel(items, list.NewDefaultDelegate(), 0, 0),
		KeyMap:       keyMap,
		saveFilePath: saveFilePath,
	}
	m.list.AdditionalShortHelpKeys = func() []key.Binding {
		return []key.Binding{
			keyMap.ToggleRead,
		}
	}
	m.list.AdditionalFullHelpKeys = func() []key.Binding {
		return []key.Binding{
			keyMap.ToggleRead,
		}
	}
	m.list.SetShowHelp(true)
	return m
}
