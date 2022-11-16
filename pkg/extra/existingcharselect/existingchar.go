package existingcharselect

import (
	"fmt"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mikejk8s/gmud/logger"
	"github.com/mikejk8s/gmud/pkg/mysqlpkg"
	"log"
)

type errMsg error

type model struct {
	spinner        spinner.Model
	quitting       bool
	err            error
	AccountOwner   string
	CharacterFound bool
}

func InitialModel(accOwner string) model {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("#0000ff"))

	return model{spinner: s, CharacterFound: false, AccountOwner: accOwner}
}

func (m model) Init() tea.Cmd {
	return m.spinner.Tick
}
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			m.quitting = true
			return m, tea.Quit
		default:
			return m, nil
		}

	case errMsg:
		m.err = msg
		return m, nil

	default:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}
}

func (m model) View() string {
	if m.err != nil {
		return m.err.Error()
	}
	str := fmt.Sprintf("\n\n   %s Loading your characters...", m.spinner.View())
	return str
}
func (m model) GetCharacterDB() {
	cDBLogger := logger.GetNewLogger()
	err := cDBLogger.AssignOutput("characterDB", "./logs/characterDBconn")
	if err != nil {
		log.Println(err)
	}
	if err != nil {
		cDBLogger.LogUtil.Errorf("Error %s connecting to characterDB during fetching the %s accounts characters: ", err, m.AccountOwner)
	}
	characters := mysqlpkg.GetCharacters(m.AccountOwner)
	fmt.Println(characters)
}
