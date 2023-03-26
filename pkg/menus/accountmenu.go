package menus

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/gliderlabs/ssh"
	"github.com/mikejk8s/gmud/pkg/characterselection/existingcharselect"
	"github.com/mikejk8s/gmud/pkg/characterselection/raceselect"
	"github.com/mikejk8s/gmud/pkg/postgrespkg"
	//"github.com/charmbracelet/wish"
)

type model struct {
	SSHSession   ssh.Session
	choices      []string         // items on the list
	cursor       int              // item our cursor is pointing at
	selected     map[int]struct{} // whats selected
	accountOwner string
}

func InitialModel(accOwner string, SSHSess ssh.Session) model {
	return model{
		SSHSession:   SSHSess, // Will not be used, just passed to the next screen.
		choices:      []string{"Play with Existing Character", "Create Character"},
		accountOwner: accOwner,
		// A map which indicates which choices are selected. We're using
		// the  map like a mathematical set. The keys refer to the indexes
		// of the `choices` slice, above.
		selected: make(map[int]struct{}),
	}
}

func (m model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit

		// The "up" and "k" keys move the cursor up
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		// The "down" and "j" keys move the cursor down
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		// The "enter" key and the spacebar (a literal space) toggle
		// the selected state for the item that the cursor is pointing at.
		case "enter", " ":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
				// Pass the account owner to the character selection screen
				// for associating the character with the account

				// Dont forget to pass SSHSession and DBConnection, a user should occupy a single session.
				dbConn := &postgrespkg.SqlConn{}
				dbConn.GetSQLConn("characters")
				switch m.choices[m.cursor] {
				case "Play with Existing Character":
					return existingcharselect.InitialModel(m.accountOwner, m.SSHSession, dbConn), nil
				case "Create Character":
					return raceselect.InitialModel(m.accountOwner, m.SSHSession, dbConn), nil
				default:
					return m, nil
				}
			}
		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m model) View() string {
	// The header
	s := "Choose your way.\n\n"
	// Iterate over our choices
	for i, choice := range m.choices {

		// Is the cursor pointing at this choice?
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor!
		}

		// Is this choice selected?
		checked := " " // not selected
		if _, ok := m.selected[i]; ok {
			checked = "x" // selected!
		}

		// Render the row
		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}
	// The footer
	s += "\nPress q to quit.\n"

	// Send the UI for rendering
	return s
}
