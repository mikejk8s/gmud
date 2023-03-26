package classelect

//
// CHARACTER SELECTION MODELS
// RACE SELECTION -> NAME SELECTION -> CLASS SELECTION (YOU ARE HERE) -> ?? (YOU ARE GOING HERE)
//
// EXISTING CHARACTER -> SELECT CHARACTER (YOU ARE HERE) -> GO TO STARTING ZONE
//
import (
	"fmt"
	"io"
	"log"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/gliderlabs/ssh"
	"github.com/mikejk8s/gmud/pkg/models"
	"github.com/mikejk8s/gmud/pkg/postgrespkg"
	"github.com/mikejk8s/gmud/pkg/zones/tutorial"
)

var (
	titleStyle        = lipgloss.NewStyle().MarginLeft(2)
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("#0099cc"))
	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
	quitTextStyle     = lipgloss.NewStyle().Margin(1, 0, 2, 4)
)

type item string

func (i item) FilterValue() string { return "" }

type itemDelegate struct{}

func (d itemDelegate) Height() int                               { return 1 }
func (d itemDelegate) Spacing() int                              { return 0 }
func (d itemDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(item)
	if !ok {
		return
	}

	str := fmt.Sprintf("%d. %s", index+1, i)

	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s string) string {
			return selectedItemStyle.Render("> " + s)
		}
	}

	_, err := fmt.Fprint(w, fn(str))
	if err != nil {
		panic(err)
	}
}

type model struct {
	DBConnection *postgrespkg.SqlConn
	SSHSession   ssh.Session
	character    *models.Character
	choiceList   list.Model
	choice       string
	cursor       int
	selected     map[int]struct{}
}

func InitialModel(characterTemp *models.Character, SSHSess ssh.Session, DBConn *postgrespkg.SqlConn) model {
	const defaultWidth = 20
	const listHeight = 14
	classes := []list.Item{
		item("Warrior"),
		item("Rogue"),
		item("Mage"),
	}
	l := list.New(classes, itemDelegate{}, defaultWidth, listHeight)
	l.Title = "Choose a class."
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle
	return model{
		DBConnection: DBConn,
		SSHSession:   SSHSess,
		choiceList:   l,
		selected:     make(map[int]struct{}),
		character:    characterTemp,
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
			// The "enter" key and the spacebar (a literal space) toggle
			// the selected state for the item that the cursor is pointing at.
		case "enter", " ":
			classChoice, ok := m.choiceList.SelectedItem().(item)
			if ok {
				m.character.Class = string(classChoice)
				//
				// this is where we insert the character into the database
				// we got:
				// account owner: from main.go
				// class: from this file
				// race: from charselection.go model
				// ID and CreatedAt is generated in charselection.go when race is picked in charselection.go model
				// Level is set to 1, alive is set to true when race is picked in charselection.go model
				//
				err := m.DBConnection.GetSQLConn("characters")
				if err != nil {
					log.Println(err)
				}
				m.DBConnection.AddCharacter(*m.character)
				m.DBConnection.Close()
				return tutorial.InitialModel(m.character, m.SSHSession), nil
			}
		}
	case tea.WindowSizeMsg:
		m.choiceList.SetWidth(msg.Width)
		return m, nil
	}
	var cmd tea.Cmd
	m.choiceList, cmd = m.choiceList.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return m.choiceList.View()
}
