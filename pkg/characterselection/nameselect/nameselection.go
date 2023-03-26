package nameselect

//
// CHARACTER SELECTION MODELS
// RACE SELECTION -> NAME SELECTION (YOU ARE HERE) -> CLASS SELECTION (YOU ARE GOING HERE)
//
// EXISTING CHARACTER -> SELECT CHARACTER (YOU ARE HERE) -> GO TO STARTING ZONE
//
import (
	"fmt"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/gliderlabs/ssh"
	"github.com/mikejk8s/gmud/pkg/characterselection/classelect"
	"github.com/mikejk8s/gmud/pkg/models"
	"github.com/mikejk8s/gmud/pkg/postgrespkg"
)

type (
	errMsg error
)
type model struct {
	DBConnection   *postgrespkg.SqlConn
	SSHSession     ssh.Session
	character      *models.Character
	input          textinput.Model
	characterClass string
	cursorMode     textinput.CursorMode
	err            error
}

func InitialModel(choice string, characterTemp *models.Character, SSHSess ssh.Session, DBConn *postgrespkg.SqlConn) model {
	ti := textinput.New()
	ti.Placeholder = "Enter here"
	ti.Focus()
	ti.CharLimit = 20
	ti.Width = 20

	return model{
		DBConnection:   DBConn,
		SSHSession:     SSHSess,
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
			return classelect.InitialModel(m.character, m.SSHSession, m.DBConnection), nil
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
