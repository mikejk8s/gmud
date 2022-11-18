package existingcharselect

import (
	"fmt"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/gliderlabs/ssh"
	"github.com/mikejk8s/gmud/logger"
	"github.com/mikejk8s/gmud/pkg/models"
	"github.com/mikejk8s/gmud/pkg/mysqlpkg"
	"github.com/mikejk8s/gmud/pkg/zones/tutorial"
	"log"
)

//
// CHARACTER SELECTION MODELS
// NEW CHARACTER -> RACE SELECTION -> NAME SELECTION -> CLASS SELECTION
//
// EXISTING CHARACTER -> SELECT CHARACTER (YOU ARE HERE) -> GO TO STARTING ZONE (YOU ARE GOING HERE)
//

var (
	appStyle          = lipgloss.NewStyle().Padding(1, 2)
	titleStyle        = lipgloss.NewStyle().MarginLeft(2)
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("#006FFF"))
	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
	quitTextStyle     = lipgloss.NewStyle().Margin(1, 0, 2, 4)
)

type item struct {
	selected         map[int]struct{}
	cursor           int
	CharacterName    string
	CharacterDetails string
}

func (i item) Title() string         { return i.CharacterName }
func (i item) Description() string   { return i.CharacterDetails }
func (i item) FilterValue() string   { return i.CharacterName }
func (i item) Style() lipgloss.Style { return itemStyle }

type model struct {
	SSHSession ssh.Session
	choiceList list.Model
	choice     string
	cursor     int
	selected   map[int]struct{}
	Character  []*models.Character
}

func InitialModel(accOwner string, SSHSess ssh.Session) model {
	// Get characters associated with the account
	tempCharacterData := GetCharacterDB(accOwner)
	var characterList []list.Item
	for i := range tempCharacterData {
		fmt.Println(tempCharacterData[i].Name)
		characterList = append(characterList, item{
			CharacterName:    tempCharacterData[i].Name,
			CharacterDetails: fmt.Sprintf("Level> %d \t Class> %s", tempCharacterData[i].Level, tempCharacterData[i].Class),
		})
	}
	var defaultWidth = 40
	var listHeight = 14
	d := list.NewDefaultDelegate()
	backgroundColor := lipgloss.Color("#000000")
	descriptionColor := lipgloss.Color("#FF9900")
	d.Styles.SelectedDesc = d.Styles.SelectedDesc.Foreground(descriptionColor).Background(backgroundColor)
	l := list.New(characterList, d, defaultWidth, listHeight)
	l.Title = "Pick your character."
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle
	return model{
		SSHSession: SSHSess,
		choiceList: l,
		selected:   make(map[int]struct{}),
		Character:  tempCharacterData,
	}
}
func (m model) Init() tea.Cmd {
	// m.GetCharacterDB()
	return nil
}
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "enter":
			_, ok := m.choiceList.SelectedItem().(item)
			if !ok {
				return m, nil
			} else {
				for i := range m.Character {
					if m.choiceList.SelectedItem().(item).CharacterName == m.Character[i].Name {
						return tutorial.InitialModel(m.Character[i], m.SSHSession), nil
					}
				}
			}
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choiceList.Items())-1 {
				m.cursor++
			}
		}
	case tea.WindowSizeMsg:
		h, v := appStyle.GetFrameSize()
		m.choiceList.SetSize(msg.Width-h, msg.Height-v)
	}
	var cmd tea.Cmd
	m.choiceList, cmd = m.choiceList.Update(msg)
	return m, cmd
}
func (m model) View() string {
	return m.choiceList.View()
}

// GetCharacterDB returns an array of characters associated with the account accOwner.
func GetCharacterDB(accOwner string) []*models.Character {
	cDBLogger := logger.GetNewLogger()
	err := cDBLogger.AssignOutput("characterDB", "./logs/characterDBconn")
	if err != nil {
		log.Println(err)
	}
	if err != nil {
		cDBLogger.LogUtil.Errorf("Error %s connecting to characterDB during fetching the %s accounts characters: ", err, accOwner)
		panic(err.Error())
	}
	return mysqlpkg.GetCharacters(accOwner)
}
