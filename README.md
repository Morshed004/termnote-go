# Termnote

A minimal terminal-based note-taking application built with Go and Bubble Tea.

Termnote lets you create, edit, save, list, and delete markdown notes directly from your terminal with a clean and fast TUI (Terminal User Interface).

---

## Features

* Create markdown notes
* Edit notes inside terminal
* Save notes instantly
* List all saved notes
* Delete notes with keyboard shortcuts
* Beautiful TUI powered by Bubble Tea
* Notes stored locally in `~/.termnote`

---

## Preview

```bash
Welcome to Termnote

┌─────────────────────────────┐
│ All Termnotes              │
├─────────────────────────────┤
│ notes.md                   │
│ ideas.md                   │
│ todo.md                    │
└─────────────────────────────┘
```

---

## Tech Stack

* Go
* Bubble Tea
* Bubbles
* Lipgloss

---

## Installation

### Clone the repository

```bash
git clone https://github.com/Morshed004/termnote-go.git
cd termnote
```

### Install dependencies

```bash
go mod tidy
```

### Run the application

```bash
go run .
```

---

## Build Binary

```bash
go build -o termnote
```

Run:

```bash
./termnote
```

---

## Keyboard Shortcuts

| Shortcut   | Action               |
| ---------- | -------------------- |
| `Ctrl + N` | Create new note      |
| `Ctrl + L` | Open notes list      |
| `Ctrl + S` | Save current note    |
| `Ctrl + D` | Delete selected note |
| `Esc`      | Close current screen |
| `Ctrl + C` | Quit application     |

---

## Notes Storage

All notes are stored locally inside:

```bash
~/.termnote
```

Each note is saved as a markdown (`.md`) file.

---

## Project Structure

```bash
termnote/
├── main.go
├── go.mod
├── go.sum
└── README.md
```

---

## How It Works

### Create a Note

Press:

```bash
Ctrl + N
```

Enter the note title and press `Enter`.

---

### Open Notes List

Press:

```bash
Ctrl + L
```

Select a note and press `Enter` to open it.

---

### Save a Note

Press:

```bash
Ctrl + S
```

---

### Delete a Note

While inside the notes list:

```bash
Ctrl + D
```

---

## Dependencies

```go
charm.land/bubbletea/v2
charm.land/bubbles/v2
charm.land/lipgloss/v2
```

---

## Future Improvements

* Search notes
* Syntax highlighting
* Auto-save
* Vim keybindings
* Folder support
* Markdown preview
* Export notes

---

## Why Termnote?

Termnote is designed for developers and terminal users who want a lightweight and distraction-free note-taking experience directly inside the terminal.