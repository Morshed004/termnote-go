package main

import (
	"fmt"
	"log"
	"os"

	"charm.land/bubbles/v2/textarea"
	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type model struct {
	NewInput               textinput.Model
	CreateFileInputVisible bool
	CurrentFile            *os.File
	Textarea               textarea.Model
	Message                string
}

var (
	vaultDir string
)

func init() {
	homeDir, err := os.UserHomeDir()

	if err != nil {
		log.Fatal("Error getting Home diractory", err)
	}

	vaultDir = fmt.Sprintf("%s/.termnote", homeDir)
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
			if m.CurrentFile != nil {
				m.CurrentFile.Close()
			}
			return m, tea.Quit

		case "ctrl+n":
			m.CreateFileInputVisible = true
			return m, nil

		case "ctrl+s":
			if m.CurrentFile != nil {

				err := os.WriteFile(
					m.CurrentFile.Name(),
					[]byte(m.Textarea.Value()),
					0644,
				)

				if err != nil {
					m.Message = "Failed to save file!"
					return m, nil
				}

				m.Message = "File saved successfully!"
			}

		case "enter":
			fileName := m.NewInput.Value()

			if fileName != "" {
				filePath := fmt.Sprintf("%s/%s.md", vaultDir, fileName)

				// Check if file already exists
				if _, err := os.Stat(filePath); err == nil {
					m.Message = "File already exists!"
					return m, nil
				}

				file, err := os.Create(filePath)
				if err != nil {
					m.Message = "Failed to create file!"
					return m, nil
				}

				m.CurrentFile = file
				m.CreateFileInputVisible = false
				m.NewInput.SetValue("")
				m.Message = "File created successfully!"
			}
		}

	}

	if m.CreateFileInputVisible {
		m.NewInput, cmd = m.NewInput.Update(msg)
	}

	if m.CurrentFile != nil {
		m.Textarea, cmd = m.Textarea.Update(msg)
	}

	return m, cmd
}

func InitializeModel() model {
	if err := os.MkdirAll(vaultDir, 0750); err != nil {
		log.Fatal(err)
	}

	ti := textinput.New()
	ti.Placeholder = "Enter your note title ..."
	ti.Focus()
	ti.CharLimit = 156
	ti.SetWidth(40)

	// Textarea

	ta := textarea.New()
	ta.Placeholder = "Write your file content ..."
	ta.Focus()
	ta.SetWidth(100)

	return model{
		NewInput:               ti,
		CreateFileInputVisible: false,
		Textarea:               ta,
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

	var MessageStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("226"))

	help := "Ctrl + C: Quit, \nCtrl + N: New File, \nCtrl + L: List, \nCtrl + S: Save, \nEsc: Back/Save\n"
	message := "Welcome to Termnote: "

	view := ""
	if m.CreateFileInputVisible {
		view = m.NewInput.View()
	}

	if m.CurrentFile != nil {
		view = m.Textarea.View()
	}

	return tea.NewView(fmt.Sprintf("\n%s\n\n%s\n\n%s\n\n%s\n", style.Render(message), view, MessageStyle.Render(m.Message), help))
}

func main() {
	p := tea.NewProgram(InitializeModel())

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
