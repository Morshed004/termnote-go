package main

import (
	"log"

	tea "charm.land/bubbletea/v2"
)

type model struct {
	message string
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func InitializeModel() model {
	return model{
		message: "🙂 Welcome to Termnote: ",
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) View() tea.View {
	return tea.NewView(m.message)
}

func main() {
	p := tea.NewProgram(InitializeModel())

	if _, err:=p.Run(); err != nil{
		log.Fatal(err)
	}
}
