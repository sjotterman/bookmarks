package ui

import (
	"fmt"
)

func (m Model) View() string {
	return fmt.Sprintf("Bookmarks:\n\n%v\n\n", m.text)
}
