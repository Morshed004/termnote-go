package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"charm.land/bubbles/v2/textarea"
	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type clearMessageMsg struct{}

type model struct {
	NewInput               textinput.Model
	CreateFileInputVisible bool
	CurrentFile            *os.File
	Textarea               textarea.Model
	Message                string
	MessageType            string // success | error
}

var (
	vaultDir string
)

func init() {
	homeDir, err := os.UserHomeDir()

	if err != nil {
		log.Fatal("Error getting Home directory", err)
	}

	vaultDir = fmt.Sprintf("%s/.termnote", homeDir)
}

func clearMessageAfter() tea.Cmd {
	return tea.Tick(2*time.Second, func(t time.Time) tea.Msg {
		return clearMessageMsg{}
	})
}

func showMessage(m model, message string, messageType string) (tea.Model, tea.Cmd) {
	m.Message = message
	m.MessageType = messageType
	return m, clearMessageAfter()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {

	case clearMessageMsg:
		m.Message = ""
		m.MessageType = ""
		return m, nil

	case tea.KeyPressMsg:

		switch msg.String() {

		case "ctrl+c":
			if m.CurrentFile != nil {
				m.CurrentFile.Close()
			}
			return m, tea.Quit

		case "ctrl+n":
			m.CreateFileInputVisible = true
			return m, nil

		case "ctrl+s":
			if m.CurrentFile == nil {
				break
			}

			if err := m.CurrentFile.Truncate(0); err != nil {
				return showMessage(m, "✗ Can't save the file", "error")
			}

			if _, err := m.CurrentFile.Seek(0, 0); err != nil {
				return showMessage(m, "✗ Can't save the file", "error")
			}

			if _, err := m.CurrentFile.WriteString(m.Textarea.Value()); err != nil {
				return showMessage(m, "✗ Can't save the file", "error")
			}

			if err := m.CurrentFile.Close(); err != nil {
				return showMessage(m, "✗ Can't save the file", "error")
			}

			m.CurrentFile = nil
			m.Textarea.SetValue("")

			return showMessage(m, "✓ File saved successfully!", "success")

		case "enter":
			fileName := m.NewInput.Value()

			if fileName != "" {
				filePath := fmt.Sprintf("%s/%s.md", vaultDir, fileName)

				// File already exists
				if _, err := os.Stat(filePath); err == nil {
					return showMessage(m, "✗ File already exists!", "error")
				}

				file, err := os.Create(filePath)
				if err != nil {
					return showMessage(m, "✗ Failed to create file!", "error")
				}

				m.CurrentFile = file
				m.CreateFileInputVisible = false
				m.NewInput.SetValue("")

				return showMessage(m, "✓ File created successfully!", "success")
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

	ta := textarea.New()
	ta.Placeholder = "Write your file content ..."
	ta.Focus()
	ta.SetWidth(100)
	ta.ShowLineNumbers = false

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
		Background(lipgloss.Color("226")).
		PaddingLeft(4).
		PaddingRight(4)

	// Dynamic message color
	messageColor := lipgloss.Color("15")

	switch m.MessageType {
	case "success":
		messageColor = lipgloss.Color("42") // green
	case "error":
		messageColor = lipgloss.Color("196") // red
	}

	messageStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(messageColor)

	help := "Ctrl + C: Quit\nCtrl + N: New File\nCtrl + L: List\nCtrl + S: Save\nEsc: Back/Save\n"

	openMessage := "Welcome to Termnote: "

	view := ""

	if m.CreateFileInputVisible {
		view = m.NewInput.View()
	}

	if m.CurrentFile != nil {
		view = m.Textarea.View()
	}

	return tea.NewView(
		fmt.Sprintf(
			"\n%s\n\n%s\n\n%s\n\n%s\n",
			style.Render(openMessage),
			view,
			messageStyle.Render(m.Message),
			help,
		),
	)
}

func main() {
	p := tea.NewProgram(InitializeModel())

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}