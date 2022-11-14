package main

import (
	"context"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/wish"
	bm "github.com/charmbracelet/wish/bubbletea"
	"github.com/felixge/fgtrace"
	mn "github.com/mikejk8s/gmud/pkg/menus"
	"github.com/mikejk8s/gmud/pkg/models"
	db "github.com/mikejk8s/gmud/pkg/mysqlpkg"
	"github.com/mikejk8s/gmud/pkg/routes"
	"github.com/muesli/termenv"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	lm "github.com/charmbracelet/wish/logging"
	"github.com/gliderlabs/ssh"
)

const (
	host = "localhost"
	port = 3131
)

func pkHandler(ctx ssh.Context, key ssh.PublicKey) bool {
	return true
}
func passHandler(ctx ssh.Context, password string) bool {
	usersDB, err := routes.ConnectUserDB()
	if err != nil {
		log.Fatalln(err)
	}
	// data := usersDB.Exec(fmt.Sprintf("SELECT password from users.users where username = '%s'", ctx.User())).First(&p)
	user := models.User{}
	rows, _ := usersDB.Query(fmt.Sprintf("SELECT password, email, name, username from users.users where username = '%s'", ctx.User()))
	for rows.Next() {
		err := rows.Scan(&user.Password, &user.Email, &user.Name, &user.Username)
		if err != nil {
			panic(err)
		}
	}
	if user.Password == password {
		credentialError := user.CheckPassword(password)
		if credentialError != nil {
			return false
		} else {
			return true
		}

	} else {
		return false
	}
}
func main() {

	// Connect to char-db mysql database and create db + tables if they don't exist
	go db.Connect()

	go func() {
		defer fgtrace.Config{Dst: fgtrace.File("fgtrace.json")}.Trace().Stop()

		http.DefaultServeMux.Handle("/debug/fgtrace", fgtrace.Config{})
		err := http.ListenAndServe(":3872", nil)
		if err != nil {
			log.Fatal(err)
		}
	}()
	// Connect to user-db mysql database and create db + tables if they don't exist
	go func() {
		_, err := routes.ConnectUserDB()
		if err != nil {
			log.Fatalln(err)
		}
	}()

	// Migrate only once
	go routes.Migration()
	/* go func() {
		_, err := tracing.JaegerTraceProvider()
		if err != nil {
			panic(err)
		}
	}() */
	// SSH server begin
	s, err := wish.NewServer(
		ssh.PasswordAuth(passHandler),
		ssh.PublicKeyAuth(pkHandler),
		wish.WithAddress(fmt.Sprintf("%s:%d", host, port)),
		wish.WithHostKeyPath(".ssh/term_info_ed25519"),
		wish.WithPublicKeyAuth(func(ctx ssh.Context, key ssh.PublicKey) bool {
			return true
		}),
		wish.WithMiddleware(
			lm.Middleware(),
			loginBubbleteaMiddleware(),
		),
	)
	if err != nil {
		log.Fatalln(err)
	}
	done := make(chan os.Signal, 0)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	log.Printf("Starting SSH server on %s:%d", host, port)
	go func() {
		if err = s.ListenAndServe(); err != nil {
			log.Fatalln(err)
		}
	}()

	<-done
	log.Println("Stopping SSH server")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer func() { cancel() }()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatalln(err)
	}
}

func loginBubbleteaMiddleware() wish.Middleware {
	login := func(m tea.Model, opts ...tea.ProgramOption) *tea.Program {
		p := tea.NewProgram(m, opts...)
		go func() {
			// if err := p.Start(); err != nil {
			// 	log.Fatalln(err)
			// }
			for {
				<-time.After(1 * time.Second)
				p.Send(timeMsg(time.Now()))
			}
		}()
		return p
	}
	teaHandler := func(s ssh.Session) *tea.Program {
		m := model{
			time:     time.Now(),
			accOwner: s.User(),
		}
		return login(m, tea.WithInput(s), tea.WithOutput(s), tea.WithAltScreen())
	}
	return bm.MiddlewareWithProgramHandler(teaHandler, termenv.ANSI256)
}

type model struct {
	time     time.Time
	accOwner string // Account owner will be used for matching the characters created from this account.
}

type timeMsg time.Time

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case timeMsg:
		m.time = time.Time(msg)
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "l", "ctrl+l":
			// return login model and make it equal to main model
			return mn.InitialModel(m.accOwner), nil // Go to the login page with passing account owner
		case "n", "ctrl+n":
			//mn.NewAccount()
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
	s := "Welcome to gmud!\n\n"
	s += "Date> " + m.time.Format(time.RFC1123) + "\n\n"
	s += "Press 'l' to go in.\n"
	return fmt.Sprintf(s)
}
