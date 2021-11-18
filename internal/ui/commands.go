package ui

import (
	"encoding/csv"
	"errors"
	"io"
	"log"
	"os"

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
		reader := csv.NewReader(file)
		items := []item{}
		for {
			records, err := reader.Read()
			if err != nil {
				if err == io.EOF {
					break
				}
				log.Fatal(err)
			}
			title := records[0]
			var isRead bool
			readStatusString := records[1]
			if readStatusString == "read" {
				isRead = true
			}
			items = append(items, item{title: title, isRead: isRead})
		}

		var updateListMsg updateBookmarkListMsg
		updateListMsg.items = items
		return updateListMsg
	}
}
