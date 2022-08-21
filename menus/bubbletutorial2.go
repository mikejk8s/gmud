// package main

// import (
// 	"fmt"
// 	"net/http"
// 	"os"
// 	"time"

// 	tea "github.com/charmbracelet/bubbletea"
// )

// const url = "https://charm.sh"

// type BubbleModel struct {
// 	status int
// 	err    error
// }

// func CheckServer() tea.Msg {

// 	// Create an HTTP client and make a GET request.
// 	c := &http.Client{Timeout: 10 * time.Second}
// 	res, err := c.Get(url)

// 	if err != nil {
// 		// There was an error making our request. Wrap the error we received
// 		// in a message and return it.
// 		return errMsg{err}
// 	}
// 	// We received a response from the server. Return the HTTP status code
// 	// as a message.
// 	return statusMsg(res.StatusCode)
// }

// type statusMsg int

// type errMsg struct{ err error }

// // For messages that contain errors its often handy to also implelement
// // the error interface on the message
// func (e errMsg) Error() string { return e.err.Error() }

// func (m model) Init() tea.Cmd {
// 	return CheckServer()
// }

// func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
// 	switch msg := msg.(type) {

// 	case statusMsg:
// 		m.status = int(msg)
// 		return m, tea.Quit

// 	case errMsg:
// 		m.err = msg.err
// 		return m, tea.Quit

// 	case tea.KeyMsg:
// 		if msg.Type == tea.KeyCtrlC {
// 			return m, tea.Quit

// 		}
// 	}
// 	return m, nil
// }

// func (m model) BubbleView() string {
// 	if m.err != nil {
// 		return fmt.Sprintf("\nWe had some trouble: %v\n\n", m.err)
// 	}

// 	s := fmt.Sprintf("checking %s ... ", url)

// 	if m.status > 0 {
// 		s += fmt.Sprintf("%d %s!", m.status, http.StatusText(m.status))
// 	}

// 	return "\n" + s + "\n\n"

// }

// func BubbleMain() {
// 	if err := tea.NewProgram(model{}).Start(); err != nil {
// 		fmt.Printf("Uh oh, there was an error: %v\n", err)
// 		os.Exit(1)

// 	}

// }

// //TODO: This is broken
