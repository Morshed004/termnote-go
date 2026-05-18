package main

import (
	"fmt"
	"log"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type model struct {
	message string
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyPressMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit

		}
	}

	return m, nil
}

func InitializeModel() model {
	return model{
		message: "Welcome to Termnote: ",
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) View() tea.View {
	var style = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("16")).
		Background(lipgloss.Color("46")).
		PaddingLeft(4).
		PaddingRight(4)
	help:= "Ctrl + C: Quit, \nCtrl + N: New File, \nCtrl + L: List, \nCtrl + S: Save, \nEsc: Back/Save\n"
	return tea.NewView(fmt.Sprintf("\n%s\n\n%s\n", style.Render(m.message), help))
}

func main() {
	p := tea.NewProgram(InitializeModel())

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
