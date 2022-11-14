package classelect

import (
	"database/sql"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/mikejk8s/gmud/pkg/models"
	"github.com/mikejk8s/gmud/pkg/mysqlpkg"
	"github.com/mikejk8s/gmud/pkg/routes"
	"log"
)

type model struct {
	character *models.Character
	choices   []string
	cursor    int
	selected  map[int]struct{}
}

func InitialModel(characterTemp *models.Character) model {
	return model{
		choices:   []string{"Warrior", "Wizard", "Thief"},
		selected:  make(map[int]struct{}),
		character: characterTemp,
	}
}

func (m model) Init() tea.Cmd {
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
				classChoice := m.choices[m.cursor]
				m.character.Class = classChoice
				usersDB, err := routes.ConnectUserDB()
				if err != nil {
					log.Println(err)
				}
				defer func(usersDB *sql.DB) {
					err := usersDB.Close()
					if err != nil {
						log.Println(err)
					}
				}(usersDB)
				mysqlpkg.AddCharacter(*m.character)
			}
		}
	}
	return m, nil
}

func (m model) View() string {
	s := "Which class would you like to login as?\n\n"

	for i, choice := range m.choices {

		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = "x"
		}

		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	s += "\nPress q to quit.\n"

	return s
}
