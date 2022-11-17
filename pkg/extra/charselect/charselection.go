package charselect

import (
	"fmt"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mikejk8s/gmud/pkg/charactersroutes"
	"github.com/mikejk8s/gmud/pkg/models"
	"io"
	"math/rand"
	"time"
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
	choiceList   list.Model       // items on the to-do list
	cursor       int              // which to-do list item our cursor is pointing at
	selected     map[int]struct{} // which to-do items are selected
	accountOwner string
}

func InitialModel(accountOwn string) model {
	const defaultWidth = 20
	const listHeight = 14
	races := []list.Item{
		item("Gandalf"),
		item("Fender"),
		item("Ghibli"),
	}
	l := list.New(races, itemDelegate{}, defaultWidth, listHeight)
	l.Title = "Choose a class."
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle
	return model{
		// Our shopping list is a grocery list
		choiceList:   l,
		accountOwner: accountOwn,
		// A map which indicates which choices are selected. We're using
		// the  map like a mathematical set. The keys refer to the indexes
		// of the `choices` slice, above.
		selected: make(map[int]struct{}),
	}
}

func (m model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
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
			raceCh, ok := m.choiceList.SelectedItem().(item)
			if ok {
				delete(m.selected, m.cursor)
				m.selected[m.cursor] = struct{}{}
				// Generate a random 5 digit string for ID
				rand.Seed(time.Now().UnixNano())
				id := rand.Intn(99999)
				// Create a new character struct
				newCharacter := models.Character{
					Race:           string(raceCh), // current selection
					ID:             id,             // random number for character identifier.
					Level:          1,              // Initial character level
					CreatedAt:      time.Now(),     // This will probably explode, change it to NOW() function while in SQL query
					Alive:          true,           // Initial character status
					CharacterOwner: m.accountOwner,
				}
				//
				// SCHEMA
				// RACE SELECTION (YOU ARE HERE) -> NAME SELECTION (YOU ARE GOING HERE) -> CLASS SELECTION
				//
				return charactersroutes.InitialModel(string(raceCh), &newCharacter), nil
			}
		}
	case tea.WindowSizeMsg:
		m.choiceList.SetWidth(msg.Width)
		return m, nil
	}

	// Return the updated model and a command to run.
	var cmd tea.Cmd
	m.choiceList, cmd = m.choiceList.Update(msg) // updates the choice list
	return m, cmd
}

func (m model) View() string {
	// shbow the choice list
	return m.choiceList.View()
}
