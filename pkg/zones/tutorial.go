package zones

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	AccountOwner   string
	CharacterFound bool
}

func InitialModel(accOwner string) model {
	return model{
		AccountOwner:   accOwner,
		CharacterFound: false,
	}
}
func (m model) Init() tea.Cmd {
	// m.GetCharacterDB()
	return nil
}
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}
func (m model) View() string {
	return fmt.Sprintf("Welcome brethen!")
}
