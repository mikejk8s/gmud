package zones

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/mikejk8s/gmud/pkg/models"
)

type model struct {
	Character *models.Character
}

func InitialModel(char *models.Character) model {
	return model{
		Character: char,
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
	return fmt.Sprintf("Welcome %s", m.Character.Name)
}
