package ui

type Model struct {
	text string
}

func NewModel() Model {
	return Model{
		text: "Hello world",
	}
}
