package ui

import (
	"bufio"
	"log"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) loadBookmarksCmd() tea.Cmd {
	return func() tea.Msg {
		saveFilePath := m.saveFilePath
		file, err := os.Open(saveFilePath)
		if err != nil {
			log.Fatal(err)
			return nil
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)

		items := []list.Item{}
		for scanner.Scan() {
			line := scanner.Text()
			items = append(items, item{title: line})
			log.Println("line", line)
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		var updateListMsg updateBookmarkListMsg
		updateListMsg.items = items
		return updateListMsg
	}
}
