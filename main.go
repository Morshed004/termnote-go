package main

import (
	"fmt"
	"log"

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type model struct {
	NewInput               textinput.Model
	CreateFileInputVisible bool
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyPressMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c":
			return m, tea.Quit

		case "ctrl+n":
			m.CreateFileInputVisible = true
			return m, nil
		}

	}

	if m.CreateFileInputVisible {
		m.NewInput, cmd = m.NewInput.Update(msg)
	}

	return m, cmd
}

func InitializeModel() model {
	ti := textinput.New()
	ti.Placeholder = "Enter your note title ..."
	ti.SetVirtualCursor(false)
	ti.Focus()
	ti.CharLimit = 156
	ti.SetWidth(40)

	return model{
		NewInput:               ti,
		CreateFileInputVisible: false,
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
	help := "Ctrl + C: Quit, \nCtrl + N: New File, \nCtrl + L: List, \nCtrl + S: Save, \nEsc: Back/Save\n"
	message := "Welcome to Termnote: "

	view := ""
	if m.CreateFileInputVisible {
		view = m.NewInput.View()
	}

	return tea.NewView(fmt.Sprintf("\n%s\n\n%s\n\n%s\n", style.Render(message), view, help))
}

func main() {
	p := tea.NewProgram(InitializeModel())

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
