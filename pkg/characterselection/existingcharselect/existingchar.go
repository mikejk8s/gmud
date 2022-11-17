package existingcharselect

import (
	"encoding/json"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/mikejk8s/gmud/logger"
	"github.com/mikejk8s/gmud/pkg/models"
	"github.com/mikejk8s/gmud/pkg/mysqlpkg"
	"log"
)

//
// SCHEMA
// NEW CHARACTER -> RACE SELECTION -> NAME SELECTION -> CLASS SELECTION
//
// EXISTING CHARACTER -> SELECT CHARACTER (YOU ARE HERE)
//
type model struct {
	Character *models.Character
}

func InitialModel(accOwner string) model {
	return model{
		Character: GetCharacterDB(accOwner),
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
	if m.Character == nil {
		return fmt.Sprintf("Loading...")
	} else {
		character, err := json.Marshal(m.Character)
		if err != nil {
			log.Println(err)
		}
		return fmt.Sprintf("Character: %s", character)
	}
}
func GetCharacterDB(accOwner string) *models.Character {
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
