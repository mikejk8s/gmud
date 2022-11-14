package charselect

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/mikejk8s/gmud/pkg/charactersroutes"
	"github.com/mikejk8s/gmud/pkg/models"
	"math/rand"
	"time"
)

type model struct {
	choices      []string         // items on the to-do list
	cursor       int              // which to-do list item our cursor is pointing at
	selected     map[int]struct{} // which to-do items are selected
	accountOwner string
}

func InitialModel(accountOwn string) model {
	return model{
		// Our shopping list is a grocery list
		choices:      []string{"Gandalf", "Fender", "Ghibli"},
		accountOwner: accountOwn,
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
				// Generate a random 5 digit string for ID
				rand.Seed(time.Now().UnixNano())
				id := rand.Intn(99999)
				// Create a new character struct
				newCharacter := models.Character{
					Race:           m.choices[m.cursor], // current selection
					ID:             id,                  // random number for character identifier.
					Level:          1,                   // Initial character level
					CreatedAt:      time.Now(),          // This will probably explode, change it to NOW() function while in SQL query
					Alive:          true,                // Initial character status
					CharacterOwner: m.accountOwner,
				}
				// Pass it to character name selection screen
				return charactersroutes.InitialModel(m.choices[m.cursor], &newCharacter), nil
			}
		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m model) View() string {
	// The header
	s := "Which character would you like to login as?\n\n"

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
