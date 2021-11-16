package ui

import (
	"bufio"
	"errors"
	"log"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) createEmptySaveFileCmd(filename string) tea.Cmd {
	return func() tea.Msg {
		emptyFile, err := os.Create(filename)
		if err != nil {
			log.Fatalf("Could not create file: %s", err)
		}
		emptyFile.Close()
		return emptyFileCreatedMsg{emptyFile.Name()}
	}
}

func (m Model) loadBookmarksCmd() tea.Cmd {
	return func() tea.Msg {
		saveFilePath := m.saveFilePath
		_, err := os.Stat(saveFilePath)
		if errors.Is(err, os.ErrNotExist) {
			return saveFileNotExistsMsg{saveFilePath}
		}
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
