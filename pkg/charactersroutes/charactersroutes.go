package charactersroutes

import (
	"fmt"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/mikejk8s/gmud/pkg/menus"
	"github.com/mikejk8s/gmud/pkg/models"
)

type (
	errMsg error
)
type model struct {
	character      *models.Character
	input          textinput.Model
	characterClass string
	cursorMode     textinput.CursorMode
	err            error
}

func InitialModel(choice string, characterTemp *models.Character) model {
	ti := textinput.New()
	ti.Placeholder = "Enter here"
	ti.Focus()
	ti.CharLimit = 20
	ti.Width = 20

	return model{
		character:      characterTemp,
		input:          ti,
		err:            nil,
		characterClass: choice,
	}
}
func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyEnter:
			m.character.Name = m.input.Value()
			return menus.InitialModel(m.character), nil
		}
	case errMsg:
		m.err = msg
		return m, nil
	}
	m.input, cmd = m.input.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return fmt.Sprintf("Enter the characters name for your %s with 20 characters being maximum:\n\n%s\n\n%s",
		m.characterClass,
		m.input.View(),
		"Ctrl+C or Esc to quit, enter to finish.",
	)
}
