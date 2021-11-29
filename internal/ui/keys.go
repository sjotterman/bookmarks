package ui

import "github.com/charmbracelet/bubbles/key"

type KeyMap struct {
	ToggleRead key.Binding
}

func newKeyMap() *KeyMap {
	return &KeyMap{
		ToggleRead: key.NewBinding(
			key.WithKeys("m"),
			key.WithHelp("m", "mark read/unread"),
		),
	}

}
