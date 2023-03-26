package combattutorial

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/gliderlabs/ssh"
	"github.com/mikejk8s/gmud/pkg/models"
)

// TUTORIAL ZONE
//         COMBAT TUTORIAL (YOU ARE HERE)
//					^
//					|
// 	???	   <- FIRST ROOM (tutorial.go) -> ?????
//				    |
//
// 	         SECOND ROOM (combattutorial.go)

/*
	// Send data to TCP server that a character has entered the tutorial zone.
	newTCP := tcpserver.TCPServer{}
	newTCP.Host = os.Getenv("TCP_HOST")
	if newTCP.Host == "" {
		newTCP.Host = "127.0.0.1"
		newTCP.Port = "8080"
	} else {
		newTCP.Port = os.Getenv("TCP_PORT")
	}
	// Server that is running on background atm.
	newTCP.NewTCPDialer()
*/

type errMsg struct {
	error
}
type model struct {
	char       *models.Character
	SSHSession ssh.Session
	textInput  textinput.Model
	err        error
}

func InitialModel(character *models.Character, SSH ssh.Session) model {
	ti := textinput.New()
	ti.Placeholder = "Pikachu"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	return model{
		char:       character,
		SSHSession: SSH,
		textInput:  ti,
		err:        nil,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.Type == tea.KeyCtrlC || msg.Type == tea.KeyEsc {
			return m, tea.Quit
		}
	case errMsg:
		m.err = msg
	default:
		m.textInput, _ = m.textInput.Update(msg)
	}

	return m, nil
}
func (m model) View() string {
	return fmt.Sprintf(
		"Send a message to your homies\n\n%s\n\n(esc to quit)\n",
		m.textInput.View(),
	)
}
