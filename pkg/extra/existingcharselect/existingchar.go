package existingcharselect

import (
	"fmt"
	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mikejk8s/gmud/logger"
	"github.com/mikejk8s/gmud/pkg/mysqlpkg"
	"log"
	"strings"
	"time"
)

type errMsg error

const (
	padding  = 2
	maxWidth = 80
)

var helpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#626262")).Render

type tickMsg time.Time

type model struct {
	percent        float64
	progress       progress.Model
	AccountOwner   string
	CharacterFound bool
}

func InitialModel(accOwner string) model {
	prog := progress.New(progress.WithScaledGradient("#FF7CCB", "#FDFF8C"))
	return model{
		progress:       prog,
		AccountOwner:   accOwner,
		CharacterFound: false,
	}
}
func tickCmd() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}
func (m model) Init() tea.Cmd {
	// m.GetCharacterDB()
	return tickCmd()
}
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		return m, tea.Quit

	case tea.WindowSizeMsg:
		m.progress.Width = msg.Width - padding*2 - 4
		if m.progress.Width > maxWidth {
			m.progress.Width = maxWidth
		}
		return m, nil

	case tickMsg:
		m.percent += 0.25
		if m.percent > 1.0 {
			m.percent = 1.0
			return m, tea.Quit
		}
		return m, tickCmd()

	default:
		return m, nil
	}
}

func (m model) View() string {
	pad := strings.Repeat(" ", padding)
	return "\n" +
		pad + m.progress.ViewAs(m.percent) + "\n\n" +
		pad + helpStyle("Press any key to quit")
}
func GetCharacterDB(accOwner string) {
	cDBLogger := logger.GetNewLogger()
	err := cDBLogger.AssignOutput("characterDB", "./logs/characterDBconn")
	if err != nil {
		log.Println(err)
	}
	if err != nil {
		cDBLogger.LogUtil.Errorf("Error %s connecting to characterDB during fetching the %s accounts characters: ", err, accOwner)
	}
	characters := mysqlpkg.GetCharacters(accOwner)
	fmt.Println(characters.Name)
}
