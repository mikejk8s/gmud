package menus

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/mikejk8s/gmud/pkg/extra/charselect"
	//"github.com/charmbracelet/wish"
)

type model struct {
	choices      []string         // items on the list
	cursor       int              // item our cursor is pointing at
	selected     map[int]struct{} // whats selected
	accountOwner string
}

func InitialModel(accOwner string) model {
	return model{
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
				switch m.choices[m.cursor] {
				case "Play with Existing Character":
				case "Create Character":
					return charselect.InitialModel(m.accountOwner), nil
				}
				return charselect.InitialModel(m.accountOwner), nil // TODO: Change this to switch case of m.choices[m.cursor]
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
